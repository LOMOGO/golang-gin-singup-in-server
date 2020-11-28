//定义user表
package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null;unique;size:64"`
	Gender string `sql:"type:ENUM('男', '女')"`
	Email string `gorm:"not null;unique;size:64"`
	Password string `gorm:"not null;unique;size:64"`
	//Avatar
}

//插入一行用户信息
func (u *User) Insert() error {
	tx := DB.Create(u)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}

//按照用户名查找
func (u *User) SelectByName() error {
	tx := DB.Where("name = ?", u.Name).First(u)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}

//按照ID查找
func (u *User)SelectByID() error {
	tx := DB.Where("id = ?", u.ID).First(u)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}

//更新密码
func UpdatePwd(id uint, newPwd string) error {
	tx := DB.Model(&User{}).Where("id = ?", id).Update("password", newPwd)
	err := tx.Error
	if err != nil {
		return err
	}
	return nil
}