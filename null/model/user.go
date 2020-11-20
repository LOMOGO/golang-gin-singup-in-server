//定义user表
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

func (u *User) Insert() error {
	tx := DB.Create(u)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Select() error {
	tx := DB.Where("name = ?", u.Name).First(u)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}