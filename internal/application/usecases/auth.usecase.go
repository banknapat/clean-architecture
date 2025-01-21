package usecases

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"clean-architecture/config"
	"clean-architecture/internal/domain/entities"
	"clean-architecture/internal/domain/repositories"
	"clean-architecture/internal/domain/value_objects"
)

type AuthUsecase interface {
	Register(creds *value_objects.Credentials) error
	Login(creds *value_objects.Credentials) (*TokenResponse, error)
	Logout(userID int) error
	RefreshToken(refreshToken string) (*TokenResponse, error)
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type authUsecase struct {
	userRepo repositories.UserRepository
	cfg      *config.Config
}

func NewAuthUsecase(userRepo repositories.UserRepository, cfg *config.Config) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (u *authUsecase) Register(creds *value_objects.Credentials) error {
	// ตรวจสอบซ้ำ username
	existingUser, err := u.userRepo.GetUserByUsername(creds.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("username already exists")
	}

	// แฮชรหัสผ่าน
	hashed, err := hashPassword(creds.Password)
	if err != nil {
		return err
	}

	user := entities.User{
		Username:     creds.Username,
		Email:        creds.Email,
		PasswordHash: hashed,
	}

	return u.userRepo.CreateUser(&user)
}

func (u *authUsecase) Login(creds *value_objects.Credentials) (*TokenResponse, error) {
	user, err := u.userRepo.GetUserByUsername(creds.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	// เช็ครหัสผ่าน
	if err := checkPasswordHash(creds.Password, user.PasswordHash); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// generate access/refresh
	accessToken, refreshToken, err := u.generateTokens(user)
	if err != nil {
		return nil, err
	}

	// บันทึก refresh token ใน db
	user.RefreshToken = refreshToken
	if err := u.userRepo.UpdateUser(user); err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *authUsecase) Logout(userID int) error {
	user, err := u.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	// clear refresh token
	user.RefreshToken = ""
	return u.userRepo.UpdateUser(user)
}

func (u *authUsecase) RefreshToken(refreshToken string) (*TokenResponse, error) {
	// หา user ที่มี refresh token ตรงกัน
	allUsers, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var foundUser *entities.User
	for _, usr := range allUsers {
		if usr.RefreshToken == refreshToken {
			foundUser = usr
			break
		}
	}

	if foundUser == nil {
		return nil, errors.New("invalid refresh token")
	}

	// Validate หรือ parse JWT refresh token (optional)
	// สมมติข้ามขั้น parse token

	// gen token ใหม่
	newAccess, newRefresh, err := u.generateTokens(foundUser)
	if err != nil {
		return nil, err
	}
	foundUser.RefreshToken = newRefresh
	err = u.userRepo.UpdateUser(foundUser)
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
	}, nil
}

func (u *authUsecase) generateTokens(user *entities.User) (string, string, error) {
	// access token
	atClaims := jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		return "", "", err
	}

	// refresh token
	rtClaims := jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(24 * time.Hour * 7).Unix(), // 7 วัน
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := rt.SignedString([]byte(u.cfg.JWTSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
