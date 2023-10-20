package routers

import (
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/smartdcs1/cdsdt/protocol-verifier-http-server/uploads"
	"gitlab.com/smartdcs1/cdsdt/protocol-verifier-http-server/downloads"
	"gitlab.com/smartdcs1/cdsdt/protocol-verifier-http-server/ioTracking"
)

func SetupRouter(
	dataPathString string,
	templates_path string,
) *gin.Engine {


	// Setup gin router
	router := gin.Default()
	templates_glob := templates_path + "/*"
	router.LoadHTMLGlob(templates_glob)
	router.MaxMultipartMemory = 8 << 20 // 8MiB

	// Health Check.
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "success", "message": "Server is healthy"})
	})

	// Upload AEReception configuration file.
	router.POST("/upload/aer-config", func(ctx *gin.Context) {
		path := dataPathString + "/aerconfig/"
		uploads.SingleUpload(path, ctx)
	})

	// Upload AEO_SVDC configuration file.
	router.POST("/upload/aeo-svdc-config", func(ctx *gin.Context) {
		path := dataPathString + "/aeo_svdc_config/"
		uploads.MultiUpload(path, ctx)
	})

	// Upload job definition file
	router.POST("/upload/job-definitions", func(ctx *gin.Context) {
		path := dataPathString + "/aeo_svdc_config/job_definitions/"
		uploads.MultiUpload(path, ctx)
	})

	// Upload Events
	router.POST("/upload/events", func(ctx *gin.Context) {
		path := dataPathString + "/events/"
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
			path := dataPathString + "/logs/verifier"
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
			path := dataPathString + "/logs/verifier"
			downloads.GetFileNames(
				path,
				"Verifier",
				ctx,
			)
		},
	)
	
	// Get aer logs
	router.POST(
		"/download/aerlog",
		func(ctx *gin.Context) {
			path := dataPathString + "/logs/reception"
			downloads.PostFileRequest(
				path,
				ctx,
			)
		},
	)

	// Get aer log filenames
	router.GET(
		"/download/aer-log-file-names",
		func(ctx *gin.Context) {
			path := dataPathString + "/logs/reception"
			downloads.GetFileNames(
				path,
				"Reception",
				ctx,
			)
		},
	)

	// POST request getting log files from specific locations
	router.POST(
		"/download/log-file-names",
		func(ctx *gin.Context) {
			downloads.HandleGetLogFileNames(
				dataPathString,
				ctx,
			)
		},
	)
	// Get log file
	router.POST(
		"/download/log-file",
		func(ctx *gin.Context) {
			downloads.HandleDownloadLogFile(
				dataPathString,
				ctx,
			)
		},
	)

	// Get number of files in aer-incoming
	router.GET(
		"/ioTracking/aer-incoming",
		func(ctx *gin.Context) {
			path := dataPathString + "/events/"
			dir, err := os.Open(path)
			if err != nil {
				ctx.String(http.StatusInternalServerError, "Request error: %s", err.Error())
			}
			ioTracking.GetNumFiles(
				dir,
				ctx,
			)
		},
	)
	
	// Get number of files in verifier-processed
	router.GET(
		"/ioTracking/verifier-processed",
		func(ctx *gin.Context) {
			path := dataPathString + "/verifier_processed"
			dir, err := os.Open(path)
			if err != nil {
				ctx.String(http.StatusInternalServerError, "Request error: %s", err.Error())
			}
			ioTracking.GetNumFiles(
				dir,
				ctx,
			)
		},
	)

	// Endpoint to clean up protocol verifier folders
	router.POST(
		"io/cleanup-test",
		func (ctx *gin.Context)  {
			dirPaths := []string {
				dataPathString + "/aeo_svdc_config/job_definitions",
				dataPathString + "/events",
				dataPathString + "/verifier_incoming",
				dataPathString + "/verifier_processed",
				dataPathString + "/invariant_store",
				dataPathString + "/job_id_store",
				dataPathString + "/logs/reception",
				dataPathString + "/logs/verifier",
			}
			suffixes := []string {
				"",
			}
			ioTracking.RestartPV(
				dirPaths,
				suffixes,
				ctx,
			)
		},
	)
	return router
}
