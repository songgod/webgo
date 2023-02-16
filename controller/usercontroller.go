package controller

import (
	"back/common"
	"back/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {

	// 获取参数
	db := common.GetDB()
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
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

	// 添加用户
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(b),
	}

	db.Create(&newUser)

	// 返回信息
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
}

func isTelephoneExist(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("telephone=?", tel).First(&user)
	return user.ID != 0
}

func Login(ctx *gin.Context) {

	db := common.GetDB()

	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "手机号码应该为11位."})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "密码不能小于6位."})
		return
	}

	// 手机号是否存在
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 442, "msg": "手机号不存在."})
		return
	}

	// 判断密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码不争取"})
		return
	}

	// 发放token到前端
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统错误"})
		log.Printf("generat token error: %v", err)
		return
	}

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}
