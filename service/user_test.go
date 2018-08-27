package service

import (
	"testing"
	"fmt"
	"github.com/go-example/model"
)

func TestGetUser(t *testing.T) {
	id := 1
	user,_ := GetUser(id)
	fmt.Println(user)
}

func TestGetAllUser(t *testing.T) {
	users, _ := GetAllUser()
	fmt.Println(users)
}

func TestCreateUser(t *testing.T) {
	var user model.User
	user.Name = "andy"
	user.Age = 21
	CreateUser(user)
	fmt.Println(user.Name,user.Age)
}
