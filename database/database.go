package database

import (
	"fmt"
	"SimpleResourceApi/config"
	"SimpleResourceApi/database/flatfile"
)

var databaseInstance Database

func InitializeDatabase() {
	database := config.Get().Database

	if database.Type == "flatfile" {
		databaseInstance = flatfile.Initialize(database.Flatfile.Location)
	}

	fmt.Println("Database initialized")
}

func Get() Database {
	return databaseInstance
}
