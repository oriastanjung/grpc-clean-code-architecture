package database

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/oriastanjung/grpc_server/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	sqlDB *sql.DB
)

func InitDB() {
	cfg := config.LoadEnv()
	database_url := cfg.DatabaseURL
	var err error

	DB, err = gorm.Open(postgres.Open(database_url), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed to connect to DB : ", err)
		panic(err)
	}
	sqlDB, err = DB.DB()
	if err != nil {
		log.Fatal("Failed to obtain db instance: ", err)
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB: ", err)
		panic(err)
	}

	// migration
	// Auto migrate all models
	MigrateDB(DB) // Call MigrateDB function to auto migrate models

}

// CloseDatabase closes the database connection
func CloseDatabase() error {
	if sqlDB != nil {
		return sqlDB.Close()
	}
	return nil
}

func GracefulShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Shutting down gracefully...")
		if err := CloseDatabase(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
		log.Println("Database connection closed")
		os.Exit(0)
	}()
}
