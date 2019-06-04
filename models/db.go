package models

import (
	"blog/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
	"time"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", conf.GetConfig().DatabaseURL)
	if err != nil {
		return nil, err
	}
	DB = db
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// 初始化表
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{}).Create(&User{
			ID:        uuid.NewV4().String(),
			UserName:  "admin",
			PassWord:  "admin",
		})
	}
	return db, err
}

// 注册'创建时间'钩子
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().UnixNano() / 1e6
		if ctf, ok := scope.FieldByName("CreatedAt"); ok {
			if ctf.IsBlank {
				ctf.Set(nowTime)
			}
		}
	}
}
