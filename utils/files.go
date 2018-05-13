package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

// DoesFileExist is a quick way to check if
// a file is already in the filesystem
func DoesFileExist(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// WriteBashScript is a quick way to write
// a bash script
func WriteBashScript(location string, content string) error {
	contents := fmt.Sprintf("#!/bin/sh\n%s", content)
	return ioutil.WriteFile(location, []byte(contents), 0755)
}
