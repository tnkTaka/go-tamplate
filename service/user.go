package service

import (
	"github.com/go-example/model"
)

func GetUser(id int) (model.User,error){
	var user model.User
	err := db.Find(&user,id).Error
	if err != nil{
		return user, err
	}

	return user,nil
}

func GetAllUser() (model.Users,error){
	var users model.Users
	err := db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users,err
}

func CreateUser(user model.User)  (model.User,error){
	err := db.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}


