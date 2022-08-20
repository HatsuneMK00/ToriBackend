package services

import (
	"ToriBackend/entity"
	"ToriBackend/global"
	"gorm.io/gorm"
)

type UserService interface {
	FindUser(id uint) (*entity.User, bool)
	FindUsersWithOffset(offset int) ([]entity.User, bool)
	AddUser(user *entity.User) (*entity.User, int64)
	DeleteUser(id uint) int64
}

type userService struct{}

func (s userService) AddUser(user *entity.User) (*entity.User, int64) {
	result := global.MysqlDB.Create(user)
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
	}
	return user, result.RowsAffected
}

func (s userService) FindUsersWithOffset(offset int) ([]entity.User, bool) {
	users := make([]entity.User, 0)
	result := global.MysqlDB.Order("id desc").Limit(10).Offset(offset).Find(&users)
	ok := true
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
		ok = false
	}
	return users, ok
}

func (s userService) FindUser(id uint) (*entity.User, bool) {
	user := entity.User{
		Model: gorm.Model{},
	}
	result := global.MysqlDB.First(&user, id)
	ok := true
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
		ok = false
	}
	return &user, ok
}

func (s userService) DeleteUser(id uint) int64 {
	result := global.MysqlDB.Delete(&entity.User{}, id)
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
	}
	return result.RowsAffected
}
