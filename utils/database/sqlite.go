package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// NewSqliteDB 初始化sqlite连接
func NewSqliteDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("waf.db"), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic(fmt.Sprintf("不能连接到数据库：%s", err.Error()))
	}
	return db
}
