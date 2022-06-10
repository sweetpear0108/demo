package model

import (
	"fmt"
	"time"
)

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Gender    int    `json:"gender"`
	Pwd       string `json:"pwd"`
	Create_ts int64  `json:"create_ts"`
	Update_ts int64  `json:"update_ts"`
}

func QueryUserById(id int) (*User, error) {
	rows := db.QueryRowx("SELECT * FROM user_reg WHERE id = $1", id)
	var user User
	err := rows.StructScan(&user)
	fmt.Println(rows)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
func Add() {
	_, err := db.Exec("INSERT INTO user_reg (id,email,name,gender,pwd,create_ts,update_ts) VALUES ($1, $2, $3, $4, $5, $6, $7)", 2, "tianwei@mxplayer.in", "tianwei",
		0, "123456", time.Now().Unix(), time.Now().Unix())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert success")
	}
}
