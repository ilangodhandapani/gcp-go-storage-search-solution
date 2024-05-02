package controllers

import (
	"APP-GO-GCP/models"
	"APP-GO-GCP/services"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateObjectMetadataController(context *gin.Context) {
	GCPobjectName := context.Param("GCPobjectName")
	GCPobjectName = strings.Replace(GCPobjectName, "/", "", 1)

	bucket := context.Param("collection")
	collection := context.Param("collection")
	projectId := context.Param("projectId")
	fileform, _ := context.FormFile("file")
	f, _ := fileform.Open()
	file, _ := ioutil.ReadAll(f)
	var m models.Metadata
	context.BindHeader(&m)
	creationTS := time.Now().String()[0:25]
	err := services.CreateObjectMetadata(bucket, file, GCPobjectName, m, creationTS)
	if err != nil {
		context.String(http.StatusExpectationFailed, fmt.Sprintf("'%s'", err))
	} else {
		docID, _, err := services.AddDocFirestore(projectId, collection, m, creationTS)
		if err != nil {
			context.String(http.StatusExpectationFailed, fmt.Sprintf("'%s'", err))
		} else {
			context.String(http.StatusCreated, fmt.Sprintf("'%s'", docID))
		}
	}
}
