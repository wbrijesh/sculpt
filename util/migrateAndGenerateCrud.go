package util

import "gorm.io/gorm"

func MigrateAndGenerateCrud(db *gorm.DB, object interface{}) (
	createFn string,
	readFn string,
	updateFn string,
	deleteFn string,
) {
	err := db.AutoMigrate(&object)
	if err != nil {
		return "", "", "", ""
	}

	createFn, _ = GenerateFunction(object, "create")
	readFn, _ = GenerateFunction(object, "read")
	updateFn, _ = GenerateFunction(object, "update")
	deleteFn, _ = GenerateFunction(object, "delete")

	return createFn, readFn, updateFn, deleteFn
}
