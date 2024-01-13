package database

import (
	"github.com/whiteyesx/go-practice/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	DSN := buildDSN("postgres", "postgres", "postgres")
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	return db
}

func buildDSN(db string, user string, pwd string) string { //TODO переменные окружения
	var res string
	res += " dbname=" + db
	res += " user=" + user
	res += " password=" + pwd
	return res
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(&model.Brand{})
	db.AutoMigrate(&model.Category{})
}
