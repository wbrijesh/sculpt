package fileUtilities

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(fileName string, path string) {
	var file *os.File
	var err error

	if path == "" {
		file, err = os.Create(fileName)
	} else {
		file, err = os.Create(path + "/" + fileName)
	}

	if err != nil {
		log.Fatal(err)
	} else {
		if path == "" {
			fmt.Println("Created file " + fileName)
		} else {
			fmt.Println("Created file " + path + "/" + fileName)
		}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
}
