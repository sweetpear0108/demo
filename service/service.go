package service

import (
	"awesomeProject/model"
	"awesomeProject/util"
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
