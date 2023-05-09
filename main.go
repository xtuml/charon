package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/smartdcs1/cdsdt/protocol-verifier-http-server/uploads"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*")
	router.MaxMultipartMemory = 8 << 20 // 8MiB

	// Health Check.
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "success", "message": "Server is healthy"})
	})

	// Upload AEReception configuration file.
	router.POST("/upload/aer-config", func(ctx *gin.Context) {
		uploads.SingleUpload("data/aerconfig/", ctx)
	})

	// Upload AEO_SVDC configuration file.
	router.POST("/upload/aeo-svdc-config", func(ctx *gin.Context) {
		uploads.MultiUpload("data/aeo_svdc_config/", ctx)
	})

	// Upload Events
	router.POST("/upload/events", func(ctx *gin.Context) {
		uploads.MultiUpload("data/events/", ctx)
	})

	// Route to web page.
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	router.Run(":9000")
}
