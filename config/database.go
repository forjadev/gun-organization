package config

import (
	"fmt"
	"os"

	"github.com/forjadev/gun-organization/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializeDatabase() (*gorm.DB, error) {
	var (
		user   = os.Getenv("DB_USER")
		passwd = os.Getenv("DB_PASSWD")
		dbname = os.Getenv("DB_NAME")
		port   = os.Getenv("DB_PORT")
	)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=localhost port=%s", user, passwd, dbname, port)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("could not open a database connection: %v", err)
	}

	err = db.AutoMigrate(&schemas.Team{}, &schemas.Member{})
	if err != nil {
		return nil, fmt.Errorf("could not automigrate: %v", err)
	}

	return db, nil
}
