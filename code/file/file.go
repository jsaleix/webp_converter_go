package file

import (
	"flag"
	"io/fs"
	"os"
	"path/filepath"
)

func GetExePath() (string, error) {
	/*	In dev mode I'm not sure where the executable will be generated,
	*	so I hardcode it to be relative to the shell
	 */
	DEVMODE := flag.Bool("dev", false, "For dev purpose it hardcodes the executable path")

	flag.Parse()

	if *DEVMODE {
		return "./", nil
	}
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	path := filepath.Dir(exePath)
	return path, nil
}

func GetFolder(path string) ([]fs.DirEntry, error) {
	content, err := os.ReadDir(path)

	if err == nil {
		return content, nil
	}
	return nil, err
}

func CreateFolder(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func FindOrCreateFolder(path string) ([]fs.DirEntry, error) {
	content, err := GetFolder(path)

	if err == nil {
		return content, nil
	}

	err = CreateFolder(path)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
