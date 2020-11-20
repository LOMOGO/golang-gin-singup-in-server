//用于明文密码的加密解密
package bcrypt

import "golang.org/x/crypto/bcrypt"

//将密码进行哈希加密
func HashPassword(pwd string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return "", err
	}
	return string(hashPwd), nil
}

//比较明文密码与哈西密码是否相同
func CompanyPassword(hashPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	return err == nil
}
