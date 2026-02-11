package routers

import (
	"bioskop-management-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/bioskop", controllers.CreateBioskop)
	router.GET("/showBioskop", controllers.ShowAllBioskop)

	return router
}
