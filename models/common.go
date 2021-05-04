package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gorm_example/config"
	"log"
)

var Db *gorm.DB

var err error

func init() {
	var err error
	dbConnectInfo := fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`,
		config.Config.DbUserName,
		config.Config.DbUserPassword,
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbName,
	)
	Db, err = gorm.Open(config.Config.DbDriverName, dbConnectInfo)

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Successfully connect database")
	}

	Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created table")
	}
}
