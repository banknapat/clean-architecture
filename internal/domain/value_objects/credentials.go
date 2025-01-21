package value_objects

type Credentials struct {
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}
