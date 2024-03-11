package main

import (
	"flag"
	"path/filepath"

	"gitlab.com/smartdcs1/cdsdt/protocol-verifier-http-server/routers"
)


func main() {
	// Define runtime flags
	dataPath := flag.String("path", "/data", "File storage directory")
	// Parse flags
	flag.Parse()
	dataPathString := *dataPath
	templates_full_path, err := filepath.Abs("./templates")
	if err != nil {
		panic(err)
	} 
	router := routers.SetupRouter(
		dataPathString,
		templates_full_path,
	)

	router.Run(":9000")

}