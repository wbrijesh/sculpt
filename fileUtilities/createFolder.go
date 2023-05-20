package fileUtilities

import (
	"fmt"
	"os"
)

func CreateFolder(folderName string, path string) {
	var err error
	if path == "" {
		err = os.Mkdir(folderName, 0755)
	} else {
		err = os.Mkdir(path+"/"+folderName, 0755)
	}

	if err != nil {
		return
	} else {
		if path == "" {
			fmt.Println("Created folder: " + folderName)
		} else {
			fmt.Println("Created folder: " + path + "/" + folderName)
		}
	}
}
