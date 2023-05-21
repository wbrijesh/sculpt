package util

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(fileName string, path string) (err error) {
	var file *os.File

	if path == "" {
		file, err = os.Create(fileName)
	} else {
		file, err = os.Create(path + "/" + fileName)
	}

	if err != nil {
		return err
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

	return nil
}
