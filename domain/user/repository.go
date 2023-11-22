package user

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

// NewUserRepository 实例化用户仓库
func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Migrate 迁移用户表
func (r *Repository) Migrate() {
	err := r.db.AutoMigrate(&User{})
	if err != nil {
		log.Printf("用户表迁移错误：%v", err)
	}
}

// InsertSampleData 添加测试数据
func (r *Repository) InsertSampleData() {
	user := NewUser("admin", "123456", "123456")
	r.db.FirstOrCreate(&user, User{UserName: user.UserName})
}

// Create 创建用户
func (r *Repository) Create(user *User) error {
	result := r.db.Create(user)
	return result.Error
}

// Update 更新用户信息
func (r *Repository) Update(user *User) error {
	return r.db.Save(user).Error
}

// GetByUserName 通过用户名查找用户
func (r *Repository) GetByUserName(username string) (User, error) {
	var user User
	err := r.db.Where("IsDeleted = ? AND UserName = ?", false, username).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// GetByUserId 通过id查找用户
func (r *Repository) GetByUserId(id uint) (User, error) {
	var user User
	err := r.db.Where("IsDeleted = ? AND ID = ?", false, id).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
