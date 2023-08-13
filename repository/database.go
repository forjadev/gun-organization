package repository

import (
	"fmt"
	"github.com/forjadev/gun-organization/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

type DatabaseService interface {
	Connect() error
	GetConnection() (*gorm.DB, error)
}

type PGDatabase struct {
	user string
	pass string
	name string
	host string
	port string
}

func NewDatabase() DatabaseService {
	return &PGDatabase{
		user: os.Getenv("DB_USER"),
		pass: os.Getenv("DB_PASS"),
		name: os.Getenv("DB_NAME"),
		host: os.Getenv("DB_HOST"),
		port: os.Getenv("DB_PORT"),
	}
}

func (d *PGDatabase) url() string {
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=localhost port=%s",
		d.user, d.pass, d.name, d.port)

	return connStr
}

func (d *PGDatabase) Connect() error {
	db, err := gorm.Open(postgres.Open(d.url()), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("could not open a database connection: %v", err)
	}

	err = db.AutoMigrate(&schemas.Team{}, &schemas.Member{})
	if err != nil {
		return fmt.Errorf("could not automigrate: %v", err)
	}

	return nil
}

func (d *PGDatabase) GetConnection() (*gorm.DB, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	return db, nil
}
