package main

import (
	controllers "APP-GO-GCP/controllers"
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/logging"
	"github.com/gin-gonic/gin"
)

var APP *gin.Engine

func getHomePageHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "appliation is up v1.0"})
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
		structuredWrite("flawless-lacing-392113")
		APP.Run("localhost:8081")
		structuredWrite("flawless-lacing-392113")
	} else {
		port := os.Getenv("PORT")
		structuredWrite("flawless-lacing-392113")
		APP.Run(":" + port)
		structuredWrite("flawless-lacing-392113")
	}

}

func structuredWrite(projectID string) {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create logging client: %v", err)
	}
	defer client.Close()
	const name = "log-example"
	logger := client.Logger(name)
	defer logger.Flush() // Ensure the entry is written.

	logger.Log(logging.Entry{
		// Log anything that can be marshaled to JSON.
		Payload: struct{ Anything string }{
			Anything: "The payload can be any type!",
		},
		Severity: logging.Info,
	})
}
