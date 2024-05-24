package main

import (
	controllers "APP-GO-GCP/controllers"
	"net/http"
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
)

var APP *gin.Engine

func getHomePageHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "appliation is up"})
}

func main() {
	APP = gin.Default()
	APP.GET("/", getHomePageHandler)
	APP.GET("/:projectId/:collection/readmetadata/:DocumentId", controllers.ReadMetadataByDocumentIdController)
	APP.GET("/:projectId/:collection/readobjectbyname/*GCPobjectName", controllers.ReadObjectByNameController)
	APP.POST("/:projectId/:collection/createobjectmetadata/*GCPobjectName", controllers.CreateObjectMetadataController)
	APP.GET("/:projectId/:collection/searchobjectbymetadata", controllers.SearchObjectByMetadataController)
	APP.DELETE("/:projectId/:collection/deleteobject/*GCPobjectName", controllers.DeleteObjectByNameController)
	if os.Getenv("LCP") == "LOCAL" {
		fmt.Println("here")
		APP.Run("localhost:8080")
	} else {
		port := os.Getenv("PORT")
		APP.Run(":" + port)

	}

}
