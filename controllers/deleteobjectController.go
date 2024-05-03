package controllers

import (
	"APP-GO-GCP/services"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeleteObjectByNameController(context *gin.Context) {
	collection := context.Param("collection")
	bucket := context.Param("collection")
	projectId := context.Param("projectId")
	GCPobjectName := context.Param("GCPobjectName")
	GCPobjectName = strings.Replace(GCPobjectName, "/", "", 1)
	err := services.DeleteObject(bucket, GCPobjectName, collection, projectId)
	if err != nil {
		context.String(http.StatusExpectationFailed, fmt.Sprintf("'%s' ", err))
		return
	} else {
		err = services.DeleteDocFirestore(projectId, collection, GCPobjectName)
		if err != nil {
			context.String(http.StatusNotFound, fmt.Sprintf("'%s' ", err))
			return
		} else {
			context.String(http.StatusOK, "File Deleted")
		}
	}
}
