package main

import (
	"github.com/joho/godotenv"
	"github.com/wbrijesh/sculpt/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
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

func WriteCrudFunctionsToFile(
	createFn string,
	readFn string,
	updateFn string,
	deleteFn string,
) {
	util.HandleError(util.CreateFolder("api", ""))
	util.HandleError(util.CreateFolder(util.StructToString(User{}), "api"))

	util.HandleError(util.CreateFile("type.go", "api/"+util.StructToString(User{})))
	util.HandleError(util.CreateFile("create.go", "api/"+util.StructToString(User{})))

	util.HandleError(util.CreateFile("read.go", "api/"+util.StructToString(User{})))
	util.HandleError(util.CreateFile("update.go", "api/"+util.StructToString(User{})))
	util.HandleError(util.CreateFile("delete.go", "api/"+util.StructToString(User{})))

	util.HandleError(util.OverwriteFile("type.go", "api/"+util.StructToString(User{}), "package "+util.StructToString(User{})+"\n"+UserTypeString))
	util.HandleError(util.OverwriteFile("create.go", "api/"+util.StructToString(User{}), createFn))
	util.HandleError(util.OverwriteFile("read.go", "api/"+util.StructToString(User{}), readFn))
	util.HandleError(util.OverwriteFile("update.go", "api/"+util.StructToString(User{}), updateFn))
	util.HandleError(util.OverwriteFile("delete.go", "api/"+util.StructToString(User{}), deleteFn))
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

	if mode == "dev" {

		//fmt.Println("Running Create Function")
		//user, status := test.RunGeneratedCreateFunction(db)
		//log.Println(user)
		//log.Println(status)

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
		WriteCrudFunctionsToFile(util.MigrateAndGenerateCrud(db, User{}))
	}
}
