package databaseUtilities

import (
	"github.com/wbrijesh/sculpt/functionGenerationUtilities"
	"gorm.io/gorm"
)

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

	createFn, _ = functionGenerationUtilities.GenerateFunction(object, "create")
	readFn, _ = functionGenerationUtilities.GenerateFunction(object, "read")
	updateFn, _ = functionGenerationUtilities.GenerateFunction(object, "update")
	deleteFn, _ = functionGenerationUtilities.GenerateFunction(object, "delete")

	return createFn, readFn, updateFn, deleteFn
}
