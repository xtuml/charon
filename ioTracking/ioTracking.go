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
	dir *os.File,
	ctx *gin.Context) {
	// Check number of files in directory
	// and send as json response with unix time
	fileNamesObjs, err := dir.Readdirnames(0)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Request error: %s", err.Error())
		return
	}
	num_files := len(fileNamesObjs)
	t := time.Now().Unix()
	jsonResponse := gin.H{
		"num_files": num_files,
		"t": t,
	}
	ctx.JSON(200, jsonResponse) 
}

func CleanFolders(
	dirPaths []string,
	suffixes []string,
) error {
	// Clean given folders of files with given suffix

	for _, dirPath := range dirPaths {
		for _, suffix := range suffixes {
			dirPathSuffix := dirPath + "/*" + suffix
			filePaths, err := filepath.Glob(
				dirPathSuffix,
			)
			if err != nil {
				return err

			}
			for _, filePath := range filePaths {
				_, fileName := filepath.Split(filePath)
				if !strings.HasPrefix(fileName, ".") {
					err := os.RemoveAll(filePath)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

type RestartPVStruct struct {
	WaitTime int64 `json:"wait_time"`
}

func CleanWaitCleanWait(
	dirPaths []string,
	suffixes []string,
	waitTime int64,
) error {
	// Function clean folders wait for a timeout and the clean and wait again
	err := CleanFolders(
		dirPaths,
		suffixes,
	)
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(waitTime) * time.Second)
	err = CleanFolders(
		dirPaths,
		suffixes,
	)
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(waitTime) * time.Second)
	return nil
}

func RestartPV(
	dirPaths []string,
	suffixes []string,
	ctx *gin.Context,
) {
	// Function to restart the Protocol Verifier
	var waitTimeObj RestartPVStruct

	// Bind the JSON Data to the fileRequest structure
	err := ctx.BindJSON(&waitTimeObj)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Request error: %s", err.Error())
	}
	err = CleanWaitCleanWait(
		dirPaths,
		suffixes,
		waitTimeObj.WaitTime,
	)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Request error: %s", err.Error())
	}
	ctx.String(http.StatusOK, "Folders Cleaned Successfully")
}