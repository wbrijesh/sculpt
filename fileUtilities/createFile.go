package fileUtilities

import (
	"log"
	"os"
)

func CreateFile(fileName string, path string) {
	file, err := os.Create(path + "/" + fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
}
