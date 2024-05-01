package database

import (
	"log"
	"os"

	"github.com/acme-corp/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	*gorm.DB
}

var DB *gorm.DB

func ConnectDb() {
	dsn := "host=" + os.Getenv("DATABASE_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=" + os.Getenv("POSTGRES_PORT") + " sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
	db.AutoMigrate(
		&models.Team{},
		&models.User{},
	)
	DB = db
}
