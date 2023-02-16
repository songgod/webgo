package controller

import (
	"back/common"
	"back/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(ctx *gin.Context) {

	db := common.GetDB()
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(name) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "名字不能为空."})
		return
	}
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "手机号码应该为11位."})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "密码不能小于6位."})
		return
	}

	if isTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "手机号已经注册."})
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}

	db.Create(&newUser)

	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
}

func isTelephoneExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("telephone=?", tel).First(&user)
	return user.ID != 0
}
