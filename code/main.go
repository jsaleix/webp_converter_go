package main

import (
	"fmt"
	"image/jpeg"
	"io/fs"
	"os"

	"golang.org/x/image/webp"

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
	content, err := os.ReadDir(path)

	if err == nil {
		return content, nil
	}

	//If folder does not exist try to create it
	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
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
	outputFolderPath := filepath.Join(exePath, OUTPUT_FOLDER_NAME)

	files, err := retrieveFolder(inputFolderPath)
	if err != nil || len(files) == 0 {
		fmt.Println("No file found")
		return
	}

	retrieveFolder(outputFolderPath)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) != ".webp" {
			continue
		}

		filePath := fmt.Sprintf("%s/%s", inputFolderPath, file.Name())
		outPath := fmt.Sprintf("%s/%s", outputFolderPath, file.Name())
		// Removing the extension from the name
		outPath = outPath[0 : len(outPath)-len(filepath.Ext(outPath))]

		f, err := os.Open(filePath)
		if err != nil {
			continue
		}
		defer f.Close()

		img, err := webp.Decode(f)
		if err != nil {
			continue
		}

		outFile, err := os.Create(outPath + ".jpeg")
		if err != nil {
			continue
		}

		defer outFile.Close()

		err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 90})
		if err != nil {
			continue
		}

		// newFile := fmt.Sprintf("%s/%s", outputFolderPath, file.Name())
		// fmt.Printf("\n%s", newFile)
		// os.Create(newFile)
	}

}
