package service

import "model"

func GetUser(id int) model.User {
	return model.GetUser(id)
}
