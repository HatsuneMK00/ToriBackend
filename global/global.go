package global

import (
	"WebAppStructure/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger  *zap.SugaredLogger
	Config  config.Server
	Viper   *viper.Viper
	MysqlDB *gorm.DB
)
