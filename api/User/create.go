package User

import (
	"fmt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return nil
}