package downloads

import (
	"os"
	"strings"
	"net/http"

	"github.com/gin-gonic/gin"
)


type fileRequest struct {
	FileName string `json:"fileName"`
}


func PostFileRequest(dirPath string, ctx *gin.Context) {
	// Add file to request and respond with 200 OK
	var fileNameObj fileRequest

	// Bind the JSON Data to the fileRequest structure
	err := ctx.BindJSON(&fileNameObj)
	if err != nil {
		ctx.String(http.StatusBadRequest, "Request error: %s", err.Error())
		return
	}
	// get file path
	filePath := dirPath + fileNameObj.FileName 
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
		if strings.Contains(fileName, fileSubString) {
			fileNames = append(fileNames, fileName)
		}
	}
	//fileNameString := "[" + strings.Join(fileNames, ",") + "]"
	jsonResponse := gin.H{"fileNames": fileNames}
	ctx.JSON(200, jsonResponse) 
}