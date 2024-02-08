package uploads

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func MultiUpload(filePath string, ctx *gin.Context) {
	// Multipart form
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.String(http.StatusBadRequest, "Request error: %s", err.Error())
		return
	}
	files := form.File["upload"]

	for _, file := range files {
		fileFull := filePath + file.Filename
		if err := ctx.SaveUploadedFile(file, fileFull); err != nil {
			ctx.String(http.StatusBadRequest, "Upload File Error: %s", err.Error())
			return
		}
	}
	ctx.String(http.StatusOK, fmt.Sprintf("'Uploaded %d Successfully", len(files)))
}

func SingleUpload(filePath string, ctx *gin.Context) {

	file, err := ctx.FormFile("upload")
	if err != nil {
		ctx.String(http.StatusBadRequest, "Request error: %s", err.Error())
		return
	}

	fileful := filePath + filepath.Base(file.Filename)
	if err := ctx.SaveUploadedFile(file, fileful); err != nil {
		ctx.String(http.StatusBadRequest, "Upload file error: `%s`", err.Error())
		return
	}

	ctx.String(http.StatusOK, "File %s uploaded successfully ", file.Filename)
}
