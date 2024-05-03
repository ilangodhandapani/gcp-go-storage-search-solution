package controllers

import (
	"fmt"
	"net/http"
	"strings"

	services "APP-GO-GCP/services"

	"github.com/gin-gonic/gin"
)

func ReadObjectByNameController(context *gin.Context) {
	GCPobjectName := context.Param("GCPobjectName")
	GCPobjectName = strings.Replace(GCPobjectName, "/", "", 1)
	bucket := context.Param("collection")
	projectId := context.Param("projectId")

	data, error := services.ReadObjectByName(projectId, bucket, GCPobjectName)
	if error != nil {
		context.String(http.StatusNotFound, fmt.Sprintf("'%s'", error))
	} else {
		context.Data(200, "application/pdf", data)
	}
}
