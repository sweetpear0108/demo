package service

import (
	"awesomeProject/model"
	"awesomeProject/util"
	"strings"
	"time"
)

func QueryById(id int) (*model.User, error) {
	rows := util.GetDB().QueryRowx("SELECT * FROM user_reg WHERE id = $1", id)
	var user model.User
	err := rows.StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Reg(email string, pwd string) (string, int, error) {
	name := strings.Split(email, "@")
	row, err := util.GetDB().Queryx("INSERT INTO user_reg (email, name, gender, pwd, create_ts, update_ts) VALUES ($1,$2,$3,$4,$5,$6) returning id",
		email, name[0], 0, pwd, time.Now().UnixNano(), time.Now().UnixNano())
	if err != nil {
		return "", 0, err
	}
	var id int
	err = row.Scan(&id)
	if err != nil {
		return "", 0, err
	}
	return name[0], id, nil
}
