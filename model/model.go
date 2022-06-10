package model

import (
	"awesomeProject/util"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func init() {

	var err error
	db, err = sqlx.Connect(util.DatabaseSetting.Type, fmt.Sprintf(fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		util.DatabaseSetting.Host,
		util.DatabaseSetting.Port,
		util.DatabaseSetting.User,
		util.DatabaseSetting.Password,
		util.DatabaseSetting.Name)))
	if err != nil {
		panic(err)
	}

}
