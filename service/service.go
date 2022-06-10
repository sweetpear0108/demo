package service

import (
	"awesomeProject/model"
)

func QueryById(id int) (*model.User, error) {
	user, err := model.QueryUserById(id)
	return user, err
}
