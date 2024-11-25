package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var CURRENT_DIRECTORY string

func Init() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf(err.Error())
	}

	currentDir, err = filepath.Abs(currentDir)
	if err != nil {
		log.Fatalf(err.Error())
	}

	CURRENT_DIRECTORY = currentDir

	if runtime.GOOS == "windows" {
		CURRENT_DIRECTORY = filepath.ToSlash(CURRENT_DIRECTORY)
	}
}
