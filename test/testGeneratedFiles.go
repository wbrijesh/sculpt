package test

type UserType struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Password  string
}

//func RunGeneratedCreateFunction(db *gorm.DB) (user *UserType, status string) {
//	user = &UserType{
//		ID:        3,
//		FirstName: "Jane",
//		LastName:  "Doe",
//		Email:     "janedoe@example.com",
//		Password:  "password",
//	}
//
//	err := User.CreateUser(db, (*User.User)(user))
//	if err != nil {
//		return user, "failed"
//	}
//
//	return user, "success"
//}
//
//func RunGeneratedReadFunction(db *gorm.DB) (user any, status string) {
//	var err error
//	user, err = User.ReadUser(db)
//	if err != nil {
//		return user, "failed"
//	}
//
//	return user, "success"
//}
//
//func RunGeneratedUpdateFunction(db *gorm.DB) (user *User.User, status string) {
//	user = (*User.User)(&UserType{
//		ID:        3,
//		FirstName: "Updated Jane",
//		LastName:  "Doe",
//		Email:     "janedoe@example.com",
//		Password:  "password",
//	})
//	var updatedUser *User.User
//	var err error
//	updatedUser, err = User.UpdateUser(db, (*User.User)(user))
//	if err != nil {
//		return updatedUser, "failed"
//	}
//
//	return updatedUser, "success"
//}
//
//func RunGeneratedDeleteFunction(db *gorm.DB) (status string) {
//	user := (*User.User)(&UserType{
//		ID:        3,
//		FirstName: "Updated Jane",
//		LastName:  "Doe",
//		Email:     "janedoe@example.com",
//		Password:  "password",
//	})
//
//	var err error
//	err = User.DeleteUser(db, (*User.User)(user))
//	if err != nil {
//		return "failed"
//	}
//
//	return "success"
//
//}
