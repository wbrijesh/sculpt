package util

import (
	"fmt"
	"os"
)

func CreateFolder(folderName string, path string) (err error) {
	if path == "" {
		err = os.Mkdir(folderName, 0755)
	} else {
		err = os.Mkdir(path+"/"+folderName, 0755)
	}

	if err != nil {
		return err
	} else {
		if path == "" {
			fmt.Println("Created folder: " + folderName)
		} else {
			fmt.Println("Created folder: " + path + "/" + folderName)
		}
	}

	return nil
}
