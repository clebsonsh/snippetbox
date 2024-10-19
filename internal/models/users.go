package models

import (
	"database/sql"
	"time"
)

type User struct {
	Created        time.Time
	Name           string
	Email          string
	HashedPassword []byte
	ID             int
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

func (m *UserModel) Authenticate(name, email, password string) error {
	return nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
