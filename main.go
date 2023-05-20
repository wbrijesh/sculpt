package main

import (
	"fmt"
	"github.com/wbrijesh/sculpt/databaseUtilities"
	"github.com/wbrijesh/sculpt/fileUtilities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"reflect"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

var UserTypeString = `
type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}
`

func StructToString(object interface{}) string {
	return reflect.TypeOf(object).Name()
}

func WriteCrudFunctionsToFile(createFn string,
	readFn string,
	updateFn string,
	deleteFn string) {

	fileUtilities.CreateFolder("api", "")
	fileUtilities.CreateFolder(StructToString(User{}), "api")

	fileUtilities.CreateFile("type.go", "api/"+StructToString(User{}))
	fileUtilities.CreateFile("create.go", "api/"+StructToString(User{}))
	fileUtilities.CreateFile("read.go", "api/"+StructToString(User{}))
	fileUtilities.CreateFile("update.go", "api/"+StructToString(User{}))
	fileUtilities.CreateFile("delete.go", "api/"+StructToString(User{}))

	fileUtilities.OverwriteFile("type.go", "api/"+StructToString(User{}), UserTypeString)
	fileUtilities.OverwriteFile("create.go", "api/"+StructToString(User{}), createFn)
	fileUtilities.OverwriteFile("read.go", "api/"+StructToString(User{}), readFn)
	fileUtilities.OverwriteFile("update.go", "api/"+StructToString(User{}), updateFn)
	fileUtilities.OverwriteFile("delete.go", "api/"+StructToString(User{}), deleteFn)
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to database")
	}

	WriteCrudFunctionsToFile(databaseUtilities.MigrateAndGenerateCrud(db, User{}))
}
