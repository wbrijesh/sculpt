package User

import "gorm.io/gorm"

func UpdateUser(db *gorm.DB, user *User) (updatedUser *User, err error) {
	userToUpdate := db.First(&User{}, user.ID)
	if userToUpdate.Error != nil {
		return user, userToUpdate.Error
	}

	result := db.Save(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
