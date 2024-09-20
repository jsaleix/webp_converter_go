package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const MODE = "DEV"

const INPUT_FOLDER_NAME = "input"
const OUTPUT_FOLDER_NAME = "output"

func getExePath() (string, error) {
	/*	In dev mode I'm not sure where the executable will be generated,
	*	so I hardcode it to be relative to the shell
	 */
	if MODE == "DEV" {
		return "./", nil
	}
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	path := filepath.Dir(exePath)
	return path, nil
}

func retrieveFolder(path string) ([]fs.DirEntry, error) {
	fmt.Printf(path)
	content, err := os.ReadDir(path)

	if err == nil {
		return content, nil
	}

	//If folder does not exist try to create it
	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
		fmt.Printf("\nOn a tout tenté")
		return nil, err
	}

	return nil, nil
}

func main() {
	exePath, err := getExePath()
	if err != nil {
		return
	}
	inputFolderPath := filepath.Join(exePath, INPUT_FOLDER_NAME)
	_, err = retrieveFolder(inputFolderPath)
	fmt.Printf("\n%s", inputFolderPath)
	if err != nil {
		fmt.Printf("\nNo file found")
	}

}
