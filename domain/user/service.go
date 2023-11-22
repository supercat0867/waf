package user

import "waf/utils/hash"

// Service 用户服务
type Service struct {
	r Repository
}

// NewUserService 实例化用户服务
func NewUserService(r Repository) *Service {
	// 迁移用户表
	r.Migrate()
	// 插入测试数据
	r.InsertSampleData()
	return &Service{
		r: r,
	}
}

// Create 创建用户
func (c *Service) Create(user *User) error {
	// 检查密码是否匹配
	if user.Password != user.Password2 {
		return ErrMismatchedPassword
	}
	// 检查用户是否存在
	_, err := c.r.GetByUserName(user.UserName)
	if err == nil {
		return ErrUserIsExist
	}
	// 创建用户
	err = c.r.Create(user)
	return err
}

// GetUserByID 通过id查找用户信息
func (c *Service) GetUserByID(userId uint) (User, error) {
	user, err := c.r.GetByUserId(userId)
	if err != nil {
		return User{}, ErrUserNotExist
	}
	return user, nil
}

// CheckUserAndPassword 检查用户是否存在并检验密码是否正确
func (c *Service) CheckUserAndPassword(username, password string) (User, error) {
	user, err := c.r.GetByUserName(username)
	if err != nil {
		return User{}, ErrUserNotExist
	}
	match := hash.CheckPassword(password, user.Password)
	if !match {
		return User{}, ErrPasswordNotCorrect
	}
	return user, nil
}

// Update 更新用户信息
func (c *Service) Update(user *User) error {
	return c.r.Update(user)
}

// ChangePassword 更改用户密码
func (c *Service) ChangePassword(user *User, oldPassword, newPassword string) error {
	// 检查密码是否匹配
	if newPassword != user.Password2 {
		return ErrMismatchedPassword
	}
	// 检查密码是否匹配
	match := hash.CheckPassword(oldPassword, user.Password)
	if !match {
		return ErrPasswordNotCorrect
	}
	// 加密新密码
	hashPassword, _ := hash.HashPassword(newPassword)
	user.Password = hashPassword
	// 更新用户信息
	return c.Update(user)
}
