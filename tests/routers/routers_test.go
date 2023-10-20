package routers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gitlab.com/smartdcs1/cdsdt/protocol-verifier-http-server/routers"
)

type jsonFileNames struct {
	FileNames []string `json:"fileNames"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func SetupServerTest() (*gin.Engine, []string) {
	dname, err := os.MkdirTemp("", "test_dir")
	check(err)
	router := routers.SetupRouter(
		dname,
		"../../templates",
	)
	dirNames := []string {
		"aeo_svdc_config/job_definitions",
		"events",
		"verifier_incoming",
		"verifier_processed",
		"invariant_store",
		"job_id_store",
		"logs/reception",
		"logs/verifier",
	}
	var filepaths []string
	for _, dirName := range dirNames {
		s := strings.Split(dirName, "/")
		dirpath := dname
		for _, dirName := range s {
			dirpath = dirpath + "/" + dirName
			err = os.Mkdir(dirpath, 0755)
			if err != nil {
				if err.Error() == "mkdir " + dirpath + ": file exists" {
					continue
				}
				check(err)
			}
		}
		f, err := os.Create(dirpath + "/file" )
		f.WriteString("test file")
		check(err)
		filepaths = append(filepaths, dirpath + "/" + f.Name())
	}
	return router, filepaths 
}


func TestRestartPV(t *testing.T) {
	router, filepaths := SetupServerTest()
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"wait_time":1}`)
	req, _ := http.NewRequest(
		"POST", 
		"/io/cleanup-test", 
		bytes.NewBuffer(jsonStr),
	)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Folders Cleaned Successfully", w.Body.String())
	var err error
	for _, filepath := range filepaths {
		_, err = os.Stat(filepath)
		if os.IsExist(err) {
			t.Errorf("test_file was not deleted")
		}
	} 
}

func TestDownloadLogFileNamesRECEPTION(t *testing.T) {
	router, _ := SetupServerTest()
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"location":"RECEPTION","file_prefix":"file"}`)
	req, _ := http.NewRequest(
		"POST", 
		"/download/log-file-names", 
		bytes.NewBuffer(jsonStr),
	)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var j jsonFileNames
    err := json.NewDecoder(w.Body).Decode(&j)
	check(err)
	assert.Contains(t, j.FileNames, "file")
}

func TestDownloadLogFileNamesVERIFIER(t *testing.T) {
	router, _ := SetupServerTest()
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"location":"VERIFIER","file_prefix":"file"}`)
	req, _ := http.NewRequest(
		"POST", 
		"/download/log-file-names", 
		bytes.NewBuffer(jsonStr),
	)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var j jsonFileNames
    err := json.NewDecoder(w.Body).Decode(&j)
	check(err)
	assert.Contains(t, j.FileNames, "file")
}

func TestDownloadLogFileNamesIncorrectLocation(t *testing.T) {
	router, _ := SetupServerTest()
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"location":"VERIFIERS","file_prefix":"file"}`)
	req, _ := http.NewRequest(
		"POST", 
		"/download/log-file-names", 
		bytes.NewBuffer(jsonStr),
	)
	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	assert.Equal(
		t,
		"Request error: the input key VERIFIERS does not exist", 
		w.Body.String(),
	)
}

func TestDownloadLogFile(t *testing.T) {
	router, _ := SetupServerTest()
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"location":"VERIFIER","fileName":"file"}`)
	req, _ := http.NewRequest(
		"POST", 
		"/download/log-file", 
		bytes.NewBuffer(jsonStr),
	)
	router.ServeHTTP(w, req)
	println(w.Body.String())
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "test file", w.Body.String())
}

func TestDownloadAerLog(t *testing.T) {
	router, _ := SetupServerTest()
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"fileName":"file"}`)
	req, _ := http.NewRequest(
		"POST", 
		"/download/aerlog", 
		bytes.NewBuffer(jsonStr),
	)
	router.ServeHTTP(w, req)
	println(w.Body.String())
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "test file", w.Body.String())
}

func TestDownloadVerifierLog(t *testing.T) {
	router, _ := SetupServerTest()
	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"fileName":"file"}`)
	req, _ := http.NewRequest(
		"POST", 
		"/download/verifierlog", 
		bytes.NewBuffer(jsonStr),
	)
	router.ServeHTTP(w, req)
	println(w.Body.String())
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "test file", w.Body.String())
}