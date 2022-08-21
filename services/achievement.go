package services

import (
	"ToriBackend/entity"
	"ToriBackend/global"
)

type AchievementService interface {
	GetAllAchievements() ([]entity.Achievement, bool)
	GetUnlockedAchievements(userId uint) ([]entity.Achievement, bool)
	CanUnlockNewAchievement(userId uint) ([]entity.Achievement, bool)
}

type achievementService struct {
}

func (s achievementService) GetAllAchievements() ([]entity.Achievement, bool) {
	var achievements []entity.Achievement
	err := global.MysqlDB.Find(&achievements).Error
	ok := true
	if err != nil {
		global.Logger.Errorf("get all achievements failed: %v", err)
		ok = false
	}
	return achievements, ok
}

func (s achievementService) GetUnlockedAchievements(userId uint) ([]entity.Achievement, bool) {
	var achievements []entity.Achievement
	err := global.MysqlDB.
		Table("users").
		Where("users.id = ?", userId).
		Joins("INNER JOIN user_achievements ua ON users.id = ua.user_id").
		Joins("INNER JOIN achievements on ua.achievement_id = achievements.id").
		Select("achievements.*").
		Find(&achievements).Error
	ok := true
	if err != nil {
		global.Logger.Errorf("get unlocked achievements failed: %v", err)
		ok = false
	}
	return achievements, ok
}

func (s achievementService) CanUnlockNewAchievement(userId uint) ([]entity.Achievement, bool) {
	//TODO implement me
	panic("implement me")
}
