package file

import (
	"io/fs"
	"os"
	"path/filepath"
	"webp_converter/config"
)

func GetExePath() (string, error) {
	/*
	*	In dev mode I'm not sure where the executable will be generated,
	*	so I hardcode it to be relative to the shell
	*	if --dev is applied
	 */

	OPTIONS := config.GetOptions()

	if OPTIONS.DevMode {
		return "./", nil
	}
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	path := filepath.Dir(exePath)
	return path, nil
}

func getFolder(path string) ([]fs.DirEntry, error) {
	content, err := os.ReadDir(path)

	if err == nil {
		return content, nil
	}
	return nil, err
}

func createFolder(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func FindOrCreateFolder(path string) ([]fs.DirEntry, error) {
	content, err := getFolder(path)

	if err == nil {
		return content, nil
	}

	err = createFolder(path)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
