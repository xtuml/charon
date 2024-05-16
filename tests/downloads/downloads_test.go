package downloads_test

import (
	"testing"
	"os"

	"github.com/stretchr/testify/assert"
	"gitlab.com/smartdcs1/cdsdt/protocol-verifier-http-server/downloads"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Tests `HandleGetLogFileLocation`
func TestHandleGetLogFileLocation(t *testing.T) {
	
	dname, err := os.MkdirTemp("", "test_dir")
	check(err)
	logsDirName := dname + "/logs"
	err = os.Mkdir(logsDirName, 0755)
	check(err)
	receptionDirName := logsDirName + "/reception"
	err = os.Mkdir(receptionDirName, 0755)
	check(err)
	verifierDirName := logsDirName + "/verifier"
	err = os.Mkdir(verifierDirName, 0755)
	check(err)
	receptionLocation, err := downloads.HandleGetLogFileLocation(
		dname,
		"RECEPTION",
	)
	check(err)
	assert.Equal(t, receptionDirName, receptionLocation)
	verifierLocation, err := downloads.HandleGetLogFileLocation(
		dname,
		"VERIFIER",
	)
	check(err)
	assert.Equal(t, verifierDirName, verifierLocation)
}

// Tests `HandleGetLogFileLocation` when an incorrect key has been provided
func TestHandleGetLogFileLocationIncorrectKey(t *testing.T) {
	receptionLocation, err := downloads.HandleGetLogFileLocation(
		"",
		"RECEPTIONS",
	)
	assert.Equal(t, "",receptionLocation)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, "the input key RECEPTIONS does not exist", err.Error())
}