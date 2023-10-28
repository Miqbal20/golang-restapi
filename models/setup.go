package models

// Import driver mysql dan gorm
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Deklarasi variable untuk mengakses DB
var DB *gorm.DB

func ConnectDB() {
	//Inisialisasi database
	dbname := "root:@tcp(localhost:3306)/go_restapi"

	// Koneksi Database
	database, err := gorm.Open(mysql.Open(dbname))

	// Panic Error
	if err != nil {
		panic(err)
	}

	// Auto Migrasi Model
	database.AutoMigrate(&Book{})
	DB = database
}
