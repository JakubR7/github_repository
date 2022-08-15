package app

import (
	"golang_mvc/git_repo/src/api/log"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	log.Info("about to map urls", "step:1", "status:pending")
	mapUrls()
	log.Info("urls succesfully mapped", "step:2", "status:completed")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
