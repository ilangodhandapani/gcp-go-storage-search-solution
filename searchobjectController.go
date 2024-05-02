package controllers

import (
	"APP-GO-GCP/models"
	"APP-GO-GCP/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchObjectByMetadataController(context *gin.Context) {
	collection := context.Param("collection")
	projectId := context.Param("projectId")

	var c models.SearchMetadata
	context.BindHeader(&c)
	if len(c.SearchAttributes) > 0 {
		results := services.SearchActiveObjects(collection, projectId, c)
		if len(results) == 0 {
			context.String(http.StatusNotFound, "No Results")
		} else {
			context.String(http.StatusFound, fmt.Sprintf("'%s'", results))
		}
	} else {
		context.String(http.StatusBadRequest, "Invalid json search string")
	}
}
