package main

import "github.com/go-example/model"

func main()  {
	db := model.GetDBConn()

	db.DropTableIfExists(&model.User{})
	db.AutoMigrate(&model.User{})
}

