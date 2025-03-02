package routes

import (
	"cloud/src/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/media/:id", controllers.GetMediaByID)
	r.POST("/media", controllers.UploadMedia)
	r.GET("/media", controllers.GetAllMedia)
}
