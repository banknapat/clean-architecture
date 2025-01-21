-- 001_init_tables.up.sql

CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    refresh_token TEXT,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS students (
    student_id SERIAL PRIMARY KEY,
    gender VARCHAR(10),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    date_of_birth DATE,
    nationality VARCHAR(50),
    ethnicity VARCHAR(50),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS teachers (
    teacher_id SERIAL PRIMARY KEY,
    gender VARCHAR(10),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    date_of_birth DATE,
    educational_qualification VARCHAR(100),
    nationality VARCHAR(50),
    ethnicity VARCHAR(50),
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS subjects (
    subject_id SERIAL PRIMARY KEY,
    subject_name VARCHAR(100) NOT NULL,
    credit_hours INT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS student_subjects (
    student_id INT NOT NULL REFERENCES students(student_id),
    subject_id INT NOT NULL REFERENCES subjects(subject_id),
    PRIMARY KEY (student_id, subject_id)
);

CREATE TABLE IF NOT EXISTS teacher_subjects (
    teacher_id INT NOT NULL REFERENCES teachers(teacher_id),
    subject_id INT NOT NULL REFERENCES subjects(subject_id),
    PRIMARY KEY (teacher_id, subject_id)
);
