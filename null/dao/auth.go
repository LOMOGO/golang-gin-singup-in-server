package dao

import "server/model"

type AuthUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (authUser *AuthUser) StoreUser() error {
	user := model.User{
		Name:     authUser.Name,
		Email:    authUser.Email,
		Password: authUser.Password,
	}
	err := user.Insert()
	if err != nil {
		return err
	}
	return nil
}

