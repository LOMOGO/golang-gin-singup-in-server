//初始化mysql数据库
package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"server/extend/conf"
	"time"
)

var DB *gorm.DB

func Setup() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DatabaseConf.User,
		conf.DatabaseConf.Password,
		conf.DatabaseConf.Host,
		conf.DatabaseConf.Port,
		conf.DatabaseConf.Dbname,
		)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("未连接到数据库")
	}
	err = DB.AutoMigrate(&User{})
	if err != nil {
		log.Print(err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxIdleTime(time.Minute * 4)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
}
