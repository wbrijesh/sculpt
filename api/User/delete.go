package User

import "gorm.io/gorm"

func DeleteUser(db *gorm.DB, user *User) (err error) {
	result := db.Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}