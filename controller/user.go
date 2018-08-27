package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/go-example/model"
	"github.com/go-example/service"
	"strconv"
)

func GetUser(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "id should be bigger than 0"})
		return
	}

	user, err := service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "no users"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetAllUser(c *gin.Context)  {
	users, err := service.GetAllUser()
	if err != nil{
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, users)
}


func CreateUser(c *gin.Context)  {
	var user model.User
	age, err := strconv.Atoi(c.PostForm("Age"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if age <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "age should be bigger than 0"})
		return
	}

	user.Name = c.PostForm("Name")
	user.Age = age

	createUser, err := service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated,createUser)
}