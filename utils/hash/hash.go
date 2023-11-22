package hash

import "golang.org/x/crypto/bcrypt"

// HashPassword 自动加盐并哈希加密
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword 检查明文密码与哈希密码是否匹配，返回布尔值
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
