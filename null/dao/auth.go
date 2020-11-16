package dao

import "server/model"

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
	Name string `json:"name" binding:"required" label:"昵称"`
	Password string `json:"password" binding:"required,max=16,min=8" label:"密码"`
}

//在数据库中搜寻该账户是否存在
func FindUserByName(name string) (pwd string, err error) {
	user := model.User{
		Name: name,
	}
	err = user.Select()
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

