package user

import (
	"gorm.io/gorm"
	"waf/utils/hash"
)

// BeforeCreate 创建用户，给明文密码加盐哈希加密
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashPassword, err := hash.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashPassword
	return nil
}
