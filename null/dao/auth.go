package dao

import (
	"gorm.io/gorm"
	"server/model"
	"time"
)

type SignupUser struct {
	Name string `json:"name" binding:"required" label:"昵称"`
	Email string `json:"email" binding:"required,email" label:"邮箱"`
	Password string `json:"password" binding:"required,max=16,min=8" label:"密码"`
	RePassword string `json:"re_password" binding:"required,max=16,min=8" label:"重复密码"`
}

//将用户存储进数据库
func (s *SignupUser) StoreUser() error {
	user := model.User{
		Name:     s.Name,
		Email:    s.Email,
		Password: s.Password,
	}
	err := user.Insert()
	if err != nil {
		return err
	}
	return nil
}

type SigninUser struct {
	Name string `json:"name" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required,max=16,min=8" label:"密码"`
}

//在数据库中搜寻该账户是否存在
func FindUserByName(name string) (pwd string, id interface{} , err error) {
	user := model.User{
		Name: name,
	}
	err = user.SelectByName()
	if err != nil {
		return "", nil, err
	}
	return user.Password, user.ID, nil
}

//利用用户ID来获取用户信息
func UseIDGetInfo(id uint) (username string, gender string, createTime time.Time, err error) {
	user := model.User{
		Model: gorm.Model{
			ID: id,
		},
	}
	err = user.SelectByID()
	if err != nil {
		return "", "", time.Time{}, err
	}
	return user.Name, user.Gender, user.CreatedAt, nil

}
