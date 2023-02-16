package router

import (
	"back/controller"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/regist", controller.Register)
	return r
}
