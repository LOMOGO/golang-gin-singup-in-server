package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null;unique;size:64"`
	Email string `gorm:"not null;unique;size:64"`
	Password string `gorm:"not null;unique;size:64"`
}

func (user *User) Insert() error {
	tx := DB.Create(user)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}