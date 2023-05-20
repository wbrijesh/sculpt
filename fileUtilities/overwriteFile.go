package fileUtilities

import (
	"fmt"
	"log"
	"os"
)

func OverwriteFile(fileName string, path string, content string) {
	var file *os.File
	var err error

	if path == "" {
		file, err = os.OpenFile(fileName, os.O_RDWR, 0644)
	} else {
		file, err = os.OpenFile(path+"/"+fileName, os.O_RDWR, 0644)
	}

	if err != nil {
		log.Fatal(err)
	} else {
		if path == "" {
			fmt.Println("Wrote to file " + fileName)
		} else {
			fmt.Println("Wrote to file " + path + "/" + fileName)
		}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = file.WriteString(content)

	if err != nil {
		log.Fatal(err)
	}
}
