package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/smartdcs1/cdsdt/protocol-verifier-http-server/uploads"
	"gitlab.com/smartdcs1/cdsdt/protocol-verifier-http-server/downloads"
)

func main() {
	// Define runtime flags
	dataPath := flag.String("path", "/data", "File storage directory")
	// Parse flags
	flag.Parse()

	// Setup gin router
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*")
	router.MaxMultipartMemory = 8 << 20 // 8MiB

	// Health Check.
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "success", "message": "Server is healthy"})
	})

	// Upload AEReception configuration file.
	router.POST("/upload/aer-config", func(ctx *gin.Context) {
		path := *dataPath + "/aerconfig/"
		uploads.SingleUpload(path, ctx)
	})

	// Upload AEO_SVDC configuration file.
	router.POST("/upload/aeo-svdc-config", func(ctx *gin.Context) {
		path := *dataPath + "/aeo_svdc_config/"
		uploads.MultiUpload(path, ctx)
	})

	// Upload job definition file
	router.POST("/upload/job-definitions", func(ctx *gin.Context) {
		path := *dataPath + "/aeo_svdc_config/job_definitions/"
		uploads.MultiUpload(path, ctx)
	})

	// Upload Events
	router.POST("/upload/events", func(ctx *gin.Context) {
		path := *dataPath + "/events/"
		uploads.MultiUpload(path, ctx)
	})

	// Route to web page.
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	// Get verifier logs
	router.POST(
		"/download/verifierlog",
		func(ctx *gin.Context) {
			path := *dataPath + "/logs/verifier/"
			downloads.PostFileRequest(
				path,
				ctx,
			)
		},
	)

	// Get all log filenames
	router.GET(
		"/download/verifier-log-file-names",
		func(ctx *gin.Context) {
			path := *dataPath + "/logs/verifier"
			downloads.GetFileNames(
				path,
				"Verifier",
				ctx,
			)
		},
	)


	router.Run(":9000")
}
