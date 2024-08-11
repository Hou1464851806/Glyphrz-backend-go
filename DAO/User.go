package DAO

import (
	"Glyphrz-go/global"
	"Glyphrz-go/model"
	"github.com/fatih/color"
	"log"
)

// GetUserListDao
// 从数据库获取用户列表
func GetUserListDao() []model.User {
	var users []model.User
	results := global.DB.Find(&users)
	if results.Error != nil {
		log.Println(results.Error)
		return []model.User{}
	}
	return users
}

func FindUserByNameDao(name string) (model.User, bool) {
	var user model.User
	result := global.DB.Where("name = ?", name).Limit(1).Find(&user)
	if result.RowsAffected == 0 {
		return model.User{}, false
	}
	return user, true
}

func AddUser(newUser *model.User) bool {
	result := global.DB.Create(newUser)
	if result.Error != nil {
		color.Red(result.Error.Error())
		return false
	}
	return true
}
