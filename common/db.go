package common

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const DBURL = "invest_dairy:invest_dairy188@tcp(great-db.mysql.ap-south-1.rds.aliyuncs.com:3306)/invest_dairy?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"

var MySQL *gorm.DB

func OpenDB() {
	var err error
	loggerI := logger.New(Mlog, logger.Config{LogLevel: logger.Info})
	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: "b_", SingularTable: true},
		Logger:         loggerI}
	MySQL, err = gorm.Open(mysql.Open(DBURL), &config)
	if err != nil {
		Mlog.Errorf("DB conection set up failed, %s\n", err.Error())
		panic(err)
	}
	db, err := MySQL.DB()
	if err != nil {
		Mlog.Errorf("get mysql db error: %s", err.Error())
		panic(err)
	}
	db.SetConnMaxLifetime(10 * time.Second)
	db.SetMaxIdleConns(200)
	db.SetMaxOpenConns(300)
	Mlog.Infof("Connecting to [%s] \n", DBURL)
}
