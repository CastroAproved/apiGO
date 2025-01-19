package config

import (
	"github.com/CastroAproved/apiGO/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})

	if err != nil{
		logger.Errorf("Sqlite opening error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil{
		logger.Errorf("Sqlite migration error: %v", err)
		return nil, err
	}

	return db, nil

}