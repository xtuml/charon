package downloads

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	RECEPTION string = "/logs/reception"
	VERIFIER string = "/logs/verifier"
)

var inputMap = map[string]string {
	"RECEPTION": RECEPTION,
	"VERIFIER": VERIFIER,
}

func GetStringToStringMap(
	key string,
	mapper map[string]string,
) (string, error) {
	value, ok := mapper[key]
	if ok {
		return value, nil
	}
	return "", errors.New("the input key " + key + " does not exist")
}


type locationFileNameRequest struct {
	Location string `json:"location" binding:"required"`
	FilePrefix string `json:"file_prefix" binding:"required"`
}


type fileRequest struct {
	FileName string `json:"fileName" binding:"required"`
	Location string `json:"location"`
}


func BindContextBodyJsonToFileRequest(
	ctx *gin.Context,
) (fileRequest, error) {
	// Add file to request
	var fileRequestObj fileRequest

	// Bind the JSON Data to the fileRequest structure
	err := ctx.BindJSON(&fileRequestObj)
	return fileRequestObj, err
}

// Handle bad requests
func HandleBadRequest(err error, ctx *gin.Context) {
	if err != nil {
		ctx.String(http.StatusBadRequest, "Request error: %s", err.Error())
		return
	}
}


// Add file to request
func PostFileRequest(dirPath string, ctx *gin.Context) {
	var fileNameObj fileRequest
	fileNameObj, err := BindContextBodyJsonToFileRequest(
		ctx,
	)
	HandleBadRequest(err, ctx)
	// get file path
	filePath := dirPath + "/" + fileNameObj.FileName 
	ctx.File(filePath)
}


func GetFileNames(
	dirPath string,
	fileSubString string,
	ctx *gin.Context) {
	// Check files in directory with given file prefix 
	// and send as json response
	fileNamesObjs, err := os.ReadDir(dirPath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Request error: %s", err.Error())
		return
	}
	fileNames := []string{} 
	for _, fileNameObj := range fileNamesObjs {
		fileName := fileNameObj.Name()
		// Filter out filenames that do not contain the given substring
		if strings.HasPrefix(fileName, fileSubString) {
			fileNames = append(fileNames, fileName)
		}
	}
	//fileNameString := "[" + strings.Join(fileNames, ",") + "]"
	jsonResponse := gin.H{"fileNames": fileNames}
	ctx.JSON(200, jsonResponse) 
}

func HandleGetLogFileLocation(
	dataPath string,
	location string,
) (string, error) {
	
	folderLocation, err := GetStringToStringMap(
		location,
		inputMap,
	)
	if err != nil {
		return "", err
	}
	dirPath := dataPath + folderLocation
	return dirPath, nil
}

func HandleGetLogFileNames(
	dataPath string,
	ctx *gin.Context,
) {
	var locationFileObj locationFileNameRequest
	// Bind the JSON Data to the fileRequest structure
	err := ctx.BindJSON(&locationFileObj)
	HandleBadRequest(err, ctx)
	dirPath, err := HandleGetLogFileLocation(
		dataPath,
		locationFileObj.Location,
	)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Request error: %s", err.Error())
		return
	}
	GetFileNames(
		dirPath,
		locationFileObj.FilePrefix,
		ctx,
	)
}

func HandleDownloadLogFile(
	dataPath string,
	ctx *gin.Context,
) {
	fileRequestObj, err := BindContextBodyJsonToFileRequest(
		ctx,
	)
	HandleBadRequest(err, ctx)
	dirPath, err := HandleGetLogFileLocation(
		dataPath,
		fileRequestObj.Location,
	)
	HandleBadRequest(err, ctx)
	filePath := dirPath + "/" + fileRequestObj.FileName
	ctx.File(filePath)
}