package user

import "time"

// User 用户模型
type User struct {
	ID        uint   `gorm:"primarykey"`
	UserName  string `gorm:"comment:登录名"`
	Password  string `gorm:"comment:密码"`
	Password2 string `gorm:"-"`
	Token     string `gorm:"comment:token"`
	IsDeleted bool
	LastLogin time.Time `gorm:"comment:上次登录时间"`
}

// NewUser 实例化用户
func NewUser(username, password, password2 string) *User {
	return &User{
		UserName:  username,
		Password:  password,
		Password2: password2,
	}
}
