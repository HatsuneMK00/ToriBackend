package core

import (
	"ToriBackend/entity"
	"ToriBackend/global"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

func InitMysqlDB() *gorm.DB {
	c := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", c.Username, c.Password, c.Path, c.Port, c.Dbname, c.Config)
	global.Logger.Infof("%s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		global.Logger.Errorf("%s", err)
		return nil
	}
	return db
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(&entity.User{}, &entity.CardPack{}, &entity.Card{}, &entity.Record{}, &entity.Course{}, &entity.UserCourse{})
	if err != nil {
		global.Logger.Errorf("Database: RegisterTables failed, err: %v", zap.Error(err))
		os.Exit(0)
	}
	err = db.SetupJoinTable(&entity.User{}, "Courses", &entity.UserCourse{})
	if err != nil {
		global.Logger.Errorf("Database: setup join table failed, err: %v", zap.Error(err))
		os.Exit(0)
	}
	global.Logger.Info("Database: Register table successfully")
}
