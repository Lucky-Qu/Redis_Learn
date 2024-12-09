package route

import (
	"Redis_Learn/03_UsingRedisAsCache/config"
	"Redis_Learn/03_UsingRedisAsCache/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/user", controller.GetUser)
	router.POST("/user", controller.CreateUser)
	return router
}

func RunRouter(r *gin.Engine) {
	err := r.Run(config.ListenAt)
	if err != nil {
		panic(err)
	}
}
