package fileUtilities

import "os"

func CreateFolder(folderName string, path string) {
	var err error
	if path == "" {
		err = os.Mkdir(folderName, 0755)
	} else {
		err = os.Mkdir(path+"/"+folderName, 0755)
	}

	if err != nil {
		return
	}
}
