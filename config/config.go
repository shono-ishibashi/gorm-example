package config

import (
	"gopkg.in/go-ini/ini.v1"
	"gorm_example/utils"
	"log"
)

type ConfigList struct {
	DbDriverName   string
	DbName         string
	DbUserName     string
	DbUserPassword string
	DbHost         string
	DbPort         string
	ServerPort     string
	LogFile        string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		DbDriverName:   cfg.Section("db").Key("db_driver_name").String(),
		DbName:         cfg.Section("db").Key("db_name").String(),
		DbUserName:     cfg.Section("db").Key("db_user_name").String(),
		DbUserPassword: cfg.Section("db").Key("db_user_password").String(),
		DbHost:         cfg.Section("db").Key("db_host").String(),
		DbPort:         cfg.Section("db").Key("db_port").String(),
		ServerPort:     cfg.Section("web").Key("port").String(),
		LogFile:        cfg.Section("web").Key("logfile").String(),
	}
}
