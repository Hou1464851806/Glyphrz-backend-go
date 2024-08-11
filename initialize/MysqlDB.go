package initialize

import (
	"Glyphrz-go/global"
	"Glyphrz-go/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlDB() {
	// 获取mysql配置
	mysqlInfo := global.Settings.MysqlInfo
	// 定义dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Name,
		mysqlInfo.Password,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.DBName,
	)
	// 打开数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 自动更新表结构
	_ = db.AutoMigrate(&model.User{})
	global.DB = db
}
