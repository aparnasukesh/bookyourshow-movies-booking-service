package sql

import (
	"fmt"
	"log"
	"sync"

	"github.com/aparnasukesh/movies-booking-svc/config"
	"github.com/aparnasukesh/movies-booking-svc/internal/app/movies"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	mutex      sync.Mutex
	isExist    map[string]bool
)

func NewSql(config config.Config) (*gorm.DB, error) {
	if dbInstance == nil && !isExist[config.DBName] {
		mutex.Lock()
		defer mutex.Unlock()
		if dbInstance == nil && !isExist[config.DBName] {
			dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Fatal(err.Error())
				return nil, err
			}
			dbInstance = db
		}
	}

	dbInstance.AutoMigrate(&movies.Movie{})

	log.Println("Successfully auto-migrated all tables.")

	return dbInstance, nil
}
