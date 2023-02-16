package main

import (
	"back/common"
	"back/router"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := common.InitDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	e := gin.Default()

	router.CollectRouter(e)

	e.Run()
}
