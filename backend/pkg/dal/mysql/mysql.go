package mysql

import (
	"fmt"

	"autograder/pkg/config"
	"autograder/pkg/model/dbm"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB(cfg *config.MysqlConfig) (*gorm.DB, error) {
	dsnPattern := "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s"
	dsn := fmt.Sprintf(dsnPattern, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Timeout)

	gormCfg := &gorm.Config{}
	if cfg.ShowSql {
		gormCfg = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	db, err := gorm.Open(mysql.Open(dsn), gormCfg)

	if err != nil {
		logrus.Errorf("[Mysql][NewDB] gorm.Open failed, error %+v", err)
		return nil, err
	}

	err = db.AutoMigrate(&dbm.User{}, &dbm.AppRunTask{})
	if err != nil {
		logrus.Errorf("[Mysql][NewDB] db.AutoMigrate failed, error %+v", err)
		return nil, err
	}

	return db, err
}
