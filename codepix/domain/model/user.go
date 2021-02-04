package model

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Base    `valid:"required"`
	Name	`valid:"required"`
	Email	`valid:"required"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func newUser(name string, email string) (*User, error) {
	user := User{
		Name: name,
		Email: Email,
	}
	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	err := user.isValid()
	if err != nil {
		return nil, err
	}
	return &user, nil
}