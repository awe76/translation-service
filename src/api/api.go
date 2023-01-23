package api

import (
	"github.com/HqOapp/translation-service/src/controllers"
	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()
	router.POST("/v2/translation", controllers.CreateTranslation)
	router.PATCH("/v2/translation", controllers.UpdateTranslation)
	router.GET("/v2/translation/search", controllers.SearchTranslation)
	router.GET("/v2/translation/:uuid", controllers.GetTranslation)
	router.DELETE("/v2/translation/:uuid", controllers.DeleteTranslation)
	router.Run()
}
