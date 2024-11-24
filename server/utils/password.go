package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword 生成密码
func GeneratePassword(password string) (string, error) {
	// 哈希密码
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hashed), err

}

// VerifyPassword 验证密码
func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == nil {
		return true
	}
	return false
}
