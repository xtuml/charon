package iotracking_test

import (
	"testing"
	"os"

	"gitlab.com/smartdcs1/cdsdt/protocol-verifier-http-server/ioTracking"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func TestCleanFolders(t *testing.T) {
	// Tests clean folders
	dname, err := os.MkdirTemp("", "test_dir")
	check(err)
	f, err := os.CreateTemp(dname, "test_file")
	check(err)
	dir_paths := []string {
		dname,
	}
	suffixes := []string {
		"",
	}
	ioTracking.CleanFolders(
		dir_paths,
		suffixes,
	)
	_, err = os.Stat(dname)
	if os.IsExist(err) {
		t.Errorf("test_dir is not in the given path")
	}
	test_file_path := dname + "/" + f.Name()
	_, err = os.Stat(test_file_path)
	if os.IsExist(err) {
		t.Errorf("test_file was not deleted")
	}
}

func TestCleanFolders_full_sub_folder(t *testing.T) {
	// Tests clean folders but with a sub folder with a file in it
	dname, err := os.MkdirTemp("", "test_dir")
	check(err)
	sub_dname, err := os.MkdirTemp(dname, "sub_test_dir")
	check(err)
	f, err := os.CreateTemp(sub_dname, "test_file")
	check(err)
	dir_paths := []string {
		dname,
	}
	suffixes := []string {
		"",
	}
	ioTracking.CleanFolders(
		dir_paths,
		suffixes,
	)
	_, err = os.Stat(dname)
	if os.IsExist(err) {
		t.Errorf("test_dir is not in the given path")
	}
	_, err = os.Stat(sub_dname)
	if os.IsExist(err) {
		t.Errorf("sub_test_dir is not in the given path")
	}
	test_file_path := sub_dname + "/" + f.Name()
	_, err = os.Stat(test_file_path)
	if os.IsExist(err) {
		t.Errorf("test_file was not deleted")
	}
}

func TestCleanWaitCleanWait(t *testing.T) {
	// Tests clean folders
	dname, err := os.MkdirTemp("", "test_dir")
	check(err)
	f, err := os.CreateTemp(dname, "test_file")
	check(err)
	dir_paths := []string {
		dname,
	}
	suffixes := []string {
		"",
	}
	waitTime := int64(1)
	ioTracking.CleanWaitCleanWait(
		dir_paths,
		suffixes,
		waitTime,
	)
	_, err = os.Stat(dname)
	if os.IsExist(err) {
		t.Errorf("test_dir is not in the given path")
	}
	test_file_path := dname + "/" + f.Name()
	_, err = os.Stat(test_file_path)
	if os.IsExist(err) {
		t.Errorf("test_file was not deleted")
	}
}