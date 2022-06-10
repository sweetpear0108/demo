package util

import (
	"awesomeProject/model"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var DatabaseSetting = &model.Database{}

var cfg *ini.File
var db *sqlx.DB

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("database", DatabaseSetting)
	db, err = sqlx.Connect(DatabaseSetting.Type, fmt.Sprintf(fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DatabaseSetting.Host,
		DatabaseSetting.Port,
		DatabaseSetting.User,
		DatabaseSetting.Password,
		DatabaseSetting.Name)))
	if err != nil {
		panic(err)
	}
}

func GetDB() *sqlx.DB {
	return db
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
