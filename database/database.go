package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB 連線到 SQLite 資料庫
func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("db.sqlite"))

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&Album{},
	)
	if err != nil {
		panic(err)
	}

	migrator := db.Migrator()
	migrator.HasTable(&Album{})

	DB = db
}
