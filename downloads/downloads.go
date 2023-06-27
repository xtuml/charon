package downloads

import (
	//"fmt"
	//"net/http"

	"github.com/gin-gonic/gin"
)


func GetFileRequest(filePath string, ctx *gin.Context) {
	// Add file to request and respond with 200 OK
	ctx.File(filePath)
}