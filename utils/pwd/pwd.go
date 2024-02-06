package pwd

import (
	"golang.org/x/crypto/bcrypt"
	"server/global"
)

// HashPwd 加密密码
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		global.Log.Error(err)
	}
	return string(hash)
}

// CheckPwd 验证密码, hash之后的密码
func CheckPwd(hashPwd string, pwd string) bool {
	byteHash := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		global.Log.Error(err)
		return false
	}
	return true
}
