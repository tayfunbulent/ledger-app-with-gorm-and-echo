package main

import (
	"log"
	"net/http"
	"ledgerApp/src/utils/routes"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"ledgerApp/src/utils/models"
)

func main() {
	dsn := "user=postgres password=password dbname=dbname host=localhost port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get the underlying sql.DB from GORM:", err)
	}
	defer sqlDB.Close()

	err = ensureAdminExists(db)
	if err != nil {
		log.Fatalf("Failed to ensure admin exists: %v", err)
	}

	router := routes.NewRouter(db)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// This function checks if there's an admin in the database when the application starts, 
// and if not, it automatically creates one.
func ensureAdminExists(db *gorm.DB) error {
	var count int64

	err := db.Model(&models.User{}).Where("role = ?", "admin").Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		adminUser := &models.User{
			Username: "admin",
			Email:    "admin@admin.com",
			Password: "admin1234",
			Role:     "admin",
		}

		return db.Create(&adminUser).Error
	}

	return nil
}
