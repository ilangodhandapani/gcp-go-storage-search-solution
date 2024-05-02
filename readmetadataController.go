package controllers

import (
	error "APP-GO-GCP/error"
	services "APP-GO-GCP/services"
	"fmt"
	"net/http"

	logging "APP-GO-GCP/logging"

	"github.com/gin-gonic/gin"
)

func ReadMetadataByDocumentIdController(context *gin.Context) {
	logging.Logging()
	DocumentId := context.Param("DocumentId")
	collection := context.Param("collection")
	projectId := context.Param("projectId")

	metadata, err := services.ReadMetadataByDocumentId(collection, projectId, DocumentId)
	if err == nil {
		context.JSON(http.StatusOK, metadata)
	} else {
		context.JSON(http.StatusNotFound, fmt.Sprintf("'%s'", error.ErrFileNotFound))
	}
}
