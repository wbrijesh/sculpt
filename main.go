package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/wbrijesh/sculpt/databaseUtilities"
	"github.com/wbrijesh/sculpt/fileUtilities"
	"github.com/wbrijesh/sculpt/test"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"reflect"
)

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Password  string
}

var UserTypeString = `
type User struct {
	ID        uint
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

	fileUtilities.OverwriteFile("type.go", "api/"+StructToString(User{}), "package "+StructToString(User{})+"\n"+UserTypeString)
	fileUtilities.OverwriteFile("create.go", "api/"+StructToString(User{}), createFn)
	fileUtilities.OverwriteFile("read.go", "api/"+StructToString(User{}), readFn)
	fileUtilities.OverwriteFile("update.go", "api/"+StructToString(User{}), updateFn)
	fileUtilities.OverwriteFile("delete.go", "api/"+StructToString(User{}), deleteFn)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	var mode string = "test"

	if mode == "test" {

		fmt.Println("Running Create Function")
		user, status := test.RunGeneratedCreateFunction(db)
		log.Println(user)
		log.Println(status)

		//fmt.Println("Running Read Function")
		//users, status := test.RunGeneratedReadFunction(db)
		//log.Println(users)
		//log.Println(status)

		//fmt.Println("Running Update Function")
		//updatedUser, status := test.RunGeneratedUpdateFunction(db)
		//log.Println(updatedUser)
		//log.Println(status)

		//fmt.Println("Running Delete Function")
		//status := test.RunGeneratedDeleteFunction(db)
		//log.Println(status)
	} else {
		WriteCrudFunctionsToFile(databaseUtilities.MigrateAndGenerateCrud(db, User{}))
	}
}
