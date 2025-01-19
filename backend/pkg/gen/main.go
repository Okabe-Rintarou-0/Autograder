package main

import (
	"fmt"

	"autograder/pkg/config"
	"autograder/pkg/model/dbm"
	"autograder/pkg/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	var (
		cfg = config.Instance.SystemDB
		err error
		db  *gorm.DB
	)
	logrus.Infof("Read mysql config: %s", utils.FormatJsonString(cfg))
	dsnPattern := "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s"
	dsn := fmt.Sprintf(dsnPattern, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Timeout)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Fatalf("Connect to db failed, error %+v", err)
		panic(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./pkg/repository/query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	// Generate default DAO interface for those specified structs
	g.ApplyBasic(dbm.User{}, dbm.AppRunTask{})

	// Execute the generator
	g.Execute()
}
