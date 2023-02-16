package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/sqlite"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(110);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
}

func main() {

	db, err := initDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	e := gin.Default()

	e.POST("/api/auth/regist", func(ctx *gin.Context) {
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

		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}

		db.Create(&newUser)

		ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "注册成功"})
	})

	e.Run()
}

func initDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "user.db")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{})

	return
}

func isTelephoneExist(db *gorm.DB, tel string) bool {
	var user User
	db.Where("telephone=?", tel).First(&user)
	return user.ID != 0
}
