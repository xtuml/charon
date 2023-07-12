package ioTracking

import (
	"os"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNumFiles(
	dirPath string,
	ctx *gin.Context) {
	// Check number of files in directory
	// and send as json response with unix time
	t := time.Now().Unix()
	fileNamesObjs, err := os.ReadDir(dirPath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Request error: %s", err.Error())
		return
	}
	num_files := len(fileNamesObjs)
	jsonResponse := gin.H{
		"num_files": num_files,
		"t": t,
	}
	ctx.JSON(200, jsonResponse) 
}
