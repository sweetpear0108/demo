package model

import (
	"fmt"
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
