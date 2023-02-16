package router

import (
	"back/controller"
	"back/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/regist", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleWare(), controller.Info)
	return r
}
