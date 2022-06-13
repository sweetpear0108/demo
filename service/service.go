package service

import (
	"awesomeProject/model"
	"awesomeProject/util"
	"database/sql"
	"log"
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

func ExistCheck(email string) bool {
	rows := util.GetDB().QueryRowx("SELECT count(*) FROM user_reg WHERE email = $1", email)
	var num int
	err := rows.Scan(&num)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	if num == 0 {
		return false
	}
	return true
}

func Reg(email string, pwd string) (string, int) {
	name := strings.Split(email, "@")
	var id int
	row := util.GetDB().QueryRowx("INSERT INTO user_reg (email, name, gender, pwd, create_ts, update_ts) VALUES ($1,$2,$3,$4,$5,$6) returning id",
		email, name[0], 0, pwd, time.Now().UnixNano(), time.Now().UnixNano())
	row.Scan(&id)
	return name[0], id
}

func Sign(email string, pwd string) (bool, int, string) {
	rows := util.GetDB().QueryRowx("SELECT * FROM user_reg WHERE email = $1 AND pwd = $2", email, pwd)
	var user model.User
	err := rows.StructScan(&user)
	if err != nil {
		return false, 0, ""
	}
	return true, user.Id, user.Name
}

func UpdatePWD(email string, pwd string) {
	util.GetDB().QueryRowx("UPDATE user_reg SET pwd = $1, update_ts = $2 WHERE email = $3",
		pwd, time.Now().UnixNano(), email)
}

func UpdateINFO(id int, name string, gender int) {
	util.GetDB().QueryRowx("UPDATE user_reg SET name = $1, update_ts = $2, gender = $3 WHERE id = $4",
		name, time.Now().UnixNano(), gender, id)
}
