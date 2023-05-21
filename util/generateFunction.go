package util

import (
	"errors"
	"reflect"
)

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

func Create` + StructToString(object) + `(db *gorm.DB, user *User) error {
	result := db.Create(user)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return nil
}`, nil
	} else if action == "read" {
		return `package ` + StructToString(object) + `

import "gorm.io/gorm"


func Read` + StructToString(object) + `(db *gorm.DB) (users []User, err error) {
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}`, nil
	} else if action == "update" {
		return `package ` + StructToString(object) + `

import "gorm.io/gorm"

func Update` + StructToString(object) + `(db *gorm.DB, user *User) (updatedUser User, err error) {
	userToUpdate := db.First(&User{}, user.ID)
	if userToUpdate.Error != nil {
		return User{}, userToUpdate.Error
	}

	result := db.Save(&user)

	if result.Error != nil {
		return User{}, result.Error
	}

	return *user, nil
}`, nil
	} else {
		return `package ` + StructToString(object) + `

import "gorm.io/gorm"

func Delete` + StructToString(object) + `(db *gorm.DB, user *User) (err error) {
	result := db.Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}`, nil
	}
}
