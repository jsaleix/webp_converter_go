package converter

import (
	"errors"
	"fmt"
	"image/jpeg"
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/image/webp"

	"webp_converter/file"
)

const INPUT_FOLDER_NAME = "input"
const OUTPUT_FOLDER_NAME = "output"

func Run() {
	// exePath, err := file.GetExePath()
	// if err != nil {
	// 	fmt.Printf("%s", err.Error())
	// 	return
	// }

	// inputFolderPath := filepath.Join(exePath, INPUT_FOLDER_NAME)
	// outputFolderPath := filepath.Join(exePath, OUTPUT_FOLDER_NAME)
	inputFolderPath, outputFolderPath, err := getPaths()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	files, err := file.FindOrCreateFolder(inputFolderPath)
	if err != nil || len(files) == 0 {
		fmt.Println("No file found")
		return
	}

	if _, err = file.FindOrCreateFolder(outputFolderPath); err != nil {
		fmt.Println("Could not find or create output folder")
		return
	}

	processed := 0

	for _, file := range files {
		err = processFile(inputFolderPath, outputFolderPath, file)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
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

func processFile(inputFolderPath, outputFolderPath string, file fs.DirEntry) error {
	if file.IsDir() || filepath.Ext(file.Name()) != ".webp" {
		return errors.New("invalid file")
	}

	filePath := fmt.Sprintf("%s/%s", inputFolderPath, file.Name())
	outPath := fmt.Sprintf("%s/%s", outputFolderPath, file.Name())
	// Removing the extension from the name
	outPath = outPath[0 : len(outPath)-len(filepath.Ext(outPath))]

	// Retrieving the file
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Retrieving the content of webp file
	img, err := webp.Decode(f)
	if err != nil {
		return err
	}

	// Creating the new jpg file
	outFile, err := os.Create(outPath + ".jpg")
	if err != nil {
		return err
	}

	defer outFile.Close()

	// Setting the content of the newly created jpg file
	err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 90})
	return err
}

func getPaths() (inputPath, outputPath string, err error) {
	exePath, err := file.GetExePath()
	if err != nil {
		return "", "", err
	}

	inputFolderPath := filepath.Join(exePath, INPUT_FOLDER_NAME)
	outputFolderPath := filepath.Join(exePath, OUTPUT_FOLDER_NAME)

	return inputFolderPath, outputFolderPath, nil
}

func GetHelp() {
	inputFolderPath, outputFolderPath, err := getPaths()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Printf("Put your webp files in: %s\n", inputFolderPath)
	fmt.Printf("Jpg files will be generated in: %s\n", outputFolderPath)
}
