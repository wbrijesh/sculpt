package User

import (
	"gorm.io/gorm"
)

func ReadUser(db *gorm.DB) (users *[]User, err error) {
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
