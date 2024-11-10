package database

import (
	"log"

	"github.com/oriastanjung/grpc_server/internal/entities"

	"gorm.io/gorm"
)

// MigrateDB migrates all database tables based on the models
func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&entities.User{}, // tambahkan semua model di sini
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	} else {
		log.Println("Database migrated successfully")
	}
}