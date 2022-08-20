package services

import (
	"ToriBackend/entity"
	"ToriBackend/global"
	"gorm.io/gorm"
	"time"
)

type RecordService interface {
	FindAllRecords(userId uint) ([]entity.Record, bool)
	FindRecordsOfMonth(userId uint, year, month int) ([]entity.Record, bool)
	AddRecord(userId uint, cardId uint, description string) (*entity.Record, int64)
	UpdateRecord(recordId uint, cardId uint, description string) int64
}

type recordService struct{}

func (s recordService) FindAllRecords(userId uint) ([]entity.Record, bool) {
	var user entity.User
	var records []entity.Record
	user.ID = userId
	err := global.MysqlDB.Model(&user).Association("Records").Find(&records)
	ok := true
	if err != nil {
		global.Logger.Errorf("%v", err)
		ok = false
	}
	return records, ok
}

func (s recordService) FindRecordsOfMonth(userId uint, year, month int) ([]entity.Record, bool) {
	var user entity.User
	var records []entity.Record
	user.ID = userId
	// create time range with year and month
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.FixedZone("Asia/Shanghai", 8*60*60))
	lastDay := firstDay.AddDate(0, 1, 0)
	global.Logger.Infof("first day: %v, last day: %v", firstDay, lastDay)
	err := global.MysqlDB.Model(&user).Association("Records").Find(&records, "created_at >= ? and created_at <= ?", firstDay, lastDay)
	ok := true
	if err != nil {
		global.Logger.Errorf("%v", err)
		ok = false
	}
	return records, ok
}

func (s recordService) AddRecord(userId uint, cardId uint, description string) (*entity.Record, int64) {
	var user entity.User
	user.ID = userId
	var card entity.Card
	card.ID = cardId
	record := entity.Record{
		Card:        card,
		Description: description,
	}
	global.MysqlDB.Create(&record)
	err := global.MysqlDB.Model(&user).Association("Records").Append(&record)
	if err != nil {
		global.Logger.Errorf("%v", err)
		return &record, 0
	}
	return &record, 1
}

func (s recordService) UpdateRecord(recordId uint, cardId uint, description string) int64 {
	record := entity.Record{
		CardID:      cardId,
		Description: description,
	}
	record.ID = recordId
	result := global.MysqlDB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&record)
	if result.Error != nil {
		global.Logger.Errorf("%v", result.Error)
		return 0
	}
	return 1
}
