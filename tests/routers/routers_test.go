package routers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/smartdcs1/cdsdt/protocol-verifier-http-server/routers"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func TestRestartPV(t *testing.T) {
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
			dirpath, err = os.MkdirTemp(dirpath, dirName)
			check(err)
		}
		check(err)
		f, err := os.CreateTemp(dirpath, "file")
		check(err)
		filepaths = append(filepaths, dirpath + "/" + f.Name())
	}
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
	for _, filepath := range filepaths {
		_, err = os.Stat(filepath)
		if os.IsExist(err) {
			t.Errorf("test_file was not deleted")
		}
	} 
}