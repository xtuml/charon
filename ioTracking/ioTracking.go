package ioTracking

import (
	"os"
	"time"
	"net/http"
	"path/filepath"
	"strings"

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

func CleanFolders(
	dirPaths []string,
	suffixes []string,
	ctx *gin.Context,
) {
	// Clean given folders of files with given suffix

	for _, dirPath := range dirPaths {
		for _, suffix := range suffixes {
			dirPathSuffix := dirPath + "/*" + suffix
			filePaths, err := filepath.Glob(
				dirPathSuffix,
			)
			if err != nil {
				ctx.String(http.StatusInternalServerError, "Request error: %s", err.Error())
				return

			}
			for _, filePath := range filePaths {
				_, fileName := filepath.Split(filePath)
				if !strings.HasPrefix(fileName, ".") {
					err := os.Remove(filePath)
					if err != nil {
						ctx.String(http.StatusInternalServerError, "Request error: %s", err.Error())
						return
					}
				}
			}
		}
	}
	ctx.String(http.StatusOK, "Folders Cleaned Successfully")
}
