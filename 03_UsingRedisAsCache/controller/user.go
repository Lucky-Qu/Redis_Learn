package controller

import (
	"Redis_Learn/03_UsingRedisAsCache/db/model"
	"Redis_Learn/03_UsingRedisAsCache/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUser(c *gin.Context) {
	var user *model.User
	var err error
	indexStr := c.Query("id")
	indexInt, err := strconv.Atoi(indexStr)
	if err != nil {
		panic(err)
	}
	user, err = service.GetUser(indexInt)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": user,
		"msg":  "success",
	})
}
func CreateUser(c *gin.Context) {
	var user *model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"data": nil,
			"msg":  err.Error(),
		})
	}
	if err := service.CreateUser(user); err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": user,
		"msg":  "success",
	})
}
