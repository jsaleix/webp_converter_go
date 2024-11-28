package converter

import (
	"errors"
	"fmt"
	"image/jpeg"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/image/webp"

	"webp_converter/config"
)

func Run() {
	inputFolderPath := config.CURRENT_DIRECTORY
	// outputFolderPath := filepath.Join(config.CURRENT_DIRECTORY, "./result")
	// err := os.Mkdir("./result", 0750)
	// if err != nil {
	// 	log.Panicf(err.Error())
	// }
	// Change of mind, the jpg files will be added next to originals
	outputFolderPath := inputFolderPath

	files, err := os.ReadDir(inputFolderPath)
	if err != nil || len(files) == 0 {
		log.Panicf("No file found")
	}

	processed := 0

	for _, file := range files {
		err = processFile(inputFolderPath, outputFolderPath, file)
		if err != nil {
			// fmt.Printf("%s\n", err.Error())
			continue
		} else {
			processed++
		}
	}

	fmt.Println("Done")
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
