package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/image/webp"
)

const INPUT_FOLDER_NAME = "input"
const OUTPUT_FOLDER_NAME = "output"

func getExePath() (string, error) {
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
		fmt.Printf("%s", err.Error())
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

	processed := 0

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".webp" {
			continue
		}

		filePath := fmt.Sprintf("%s/%s", inputFolderPath, file.Name())
		outPath := fmt.Sprintf("%s/%s", outputFolderPath, file.Name())
		// Removing the extension from the name
		outPath = outPath[0 : len(outPath)-len(filepath.Ext(outPath))]

		// Retrieving the file
		f, err := os.Open(filePath)
		if err != nil {
			continue
		}
		defer f.Close()

		// Retrieving the content of webp file
		img, err := webp.Decode(f)
		if err != nil {
			continue
		}

		// Creating the new jpg file
		outFile, err := os.Create(outPath + ".jpg")
		if err != nil {
			continue
		}

		defer outFile.Close()

		// Setting the content of the newly created jpg file
		err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 90})
		if err != nil {
			continue
		} else {
			processed++
		}
	}

	fmt.Printf("Done\n")
	fmt.Printf("Processed files: %d\n", processed)
	if processed > 0 {
		fmt.Printf("Output dir. %s\n", outputFolderPath)
	}

}
