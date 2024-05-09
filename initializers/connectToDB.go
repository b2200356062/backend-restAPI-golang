package initializers

import (
	"staj/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// global database variable
var DB *gorm.DB

func ConnectToDB() {
	var err error

	localDB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// without this conversion global database is not being initialized
	DB = localDB
	if DB == nil {
		panic("failed to initialize global db")
	}

	// auto migration for create, update and delete times
	DB.AutoMigrate(&models.User{}, &models.ToDoList{}, &models.ToDoListMessage{})

}
