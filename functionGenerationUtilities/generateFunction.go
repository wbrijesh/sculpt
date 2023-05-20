package functionGenerationUtilities

import (
	"errors"
	"reflect"
)

func StructToString(object interface{}) string {
	return reflect.TypeOf(object).Name()
}

func GenerateFunction(object interface{}, action string) (string, error) {
	if reflect.TypeOf(object).Kind() != reflect.Struct {
		return `func(db *gorm.DB, object interface{}) {}`,
			errors.New("object passed to generateFunction is not a struct")
	}

	if action != "create" && action != "read" && action != "update" && action != "delete" {
		return `func(db *gorm.DB, object interface{}) {}`,
			errors.New("action passed to generateFunction is not a valid action")
	}

	if action == "create" {
		return `package ` + StructToString(object) + `

import (
	"fmt"
	"gorm.io/gorm"
)

func Create` + StructToString(object) + `(db *gorm.DB, object interface{}) {
	result := db.Create(&object)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println(result.RowsAffected)
}`, nil
	} else if action == "read" {
		return `package ` + StructToString(object) + `

import (
	"fmt"
	"gorm.io/gorm"
)

func Read` + StructToString(object) + `(db *gorm.DB, object interface{}) {

			result := db.Find(&objects)

			if result.Error != nil {
				fmt.Println(result.Error)
			}

			fmt.Println(result.RowsAffected)
		}`, nil
	} else if action == "update" {
		return `package ` + StructToString(object) + `

import (
	"fmt"
	"gorm.io/gorm"
)

func Update` + StructToString(object) + `(db *gorm.DB, object interface{}) {
			result := db.Save(&object)

			if result.Error != nil {
				fmt.Println(result.Error)
			}

			fmt.Println(result.RowsAffected)
		}`, nil
	} else {
		return `package ` + StructToString(object) + `

import (
	"fmt"
	"gorm.io/gorm"
)

func Delete` + StructToString(object) + `(db *gorm.DB, object interface{}) {
			result := db.Delete(&object)

			if result.Error != nil {
				fmt.Println(result.Error)
			}

			fmt.Println(result.RowsAffected)
		}`, nil
	}
}
