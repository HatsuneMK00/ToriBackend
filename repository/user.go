package repository

import (
	"ToriBackend/entity"
	"ToriBackend/global"
	"gorm.io/gorm"
)

type UserRepository interface {
	// AddUser Add a user
	AddUser(user *entity.User) (*entity.User, int64)
	// AddUsers Add multiple users at once
	AddUsers(users []entity.User) ([]entity.User, int64)
	// FindUser Find a user by its id
	FindUser(id uint) (*entity.User, bool)
	// FindUsersWithOffset Return 10 records ordered by primary key ID, desc. Return full content of entity.User
	FindUsersWithOffset(offset int) ([]entity.User, bool)
	// FindUsersSummaryWithOffset Return 10 records ordered by primary key ID, desc. Only return a subset of entity.User
	FindUsersSummaryWithOffset(offset int) ([]entity.UserSummary, bool)
	// UpdateUserInfo Update user info. Only update non-zero value
	UpdateUserInfo(user *entity.User) (*entity.User, int64)
	// DeleteUser Delete user specified by id
	DeleteUser(id uint) int64
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo() UserRepository {
	repo := userRepository{
		db: global.MysqlDB,
	}
	return &repo
}

func (m *userRepository) AddUser(user *entity.User) (*entity.User, int64) {
	result := m.db.Create(user)
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
	}
	return user, result.RowsAffected
}

func (m *userRepository) AddUsers(users []entity.User) ([]entity.User, int64) {
	result := m.db.Create(&users)
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
	}
	return users, result.RowsAffected
}

func (m *userRepository) FindUser(id uint) (*entity.User, bool) {
	user := entity.User{
		Model: gorm.Model{},
	}
	result := m.db.First(&user, id)
	ok := true
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
		ok = false
	}
	return &user, ok
}

func (m *userRepository) FindUsersWithOffset(offset int) ([]entity.User, bool) {
	users := make([]entity.User, 0)
	result := m.db.Order("id desc").Limit(10).Offset(offset).Find(&users)
	ok := true
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
		ok = false
	}
	return users, ok
}

func (m *userRepository) FindUsersSummaryWithOffset(offset int) ([]entity.UserSummary, bool) {
	shortUsers := make([]entity.UserSummary, 0)
	result := m.db.Model(&entity.User{}).Order("id desc").Limit(10).Offset(offset).Find(&shortUsers)
	ok := true
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
		ok = false
	}
	return shortUsers, ok
}

func (m *userRepository) UpdateUserInfo(user *entity.User) (*entity.User, int64) {
	result := m.db.Model(&user).Updates(&user)
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
	}
	return user, result.RowsAffected
}

func (m *userRepository) DeleteUser(id uint) int64 {
	result := m.db.Delete(&entity.User{}, id)
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
	}
	return result.RowsAffected
}
