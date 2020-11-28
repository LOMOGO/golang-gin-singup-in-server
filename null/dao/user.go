package dao

import (
	"server/extend/utils"
	"server/model"
)

//更新用户密码
func UpdateUserPwd(id uint, newPwd string) error {
	newPwd = utils.MakeSha1(newPwd)
	err := model.UpdatePwd(id, newPwd)
	if err != nil {
		return err
	}
	return nil
}
