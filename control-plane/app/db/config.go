package db

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Model represents the base model that all models should embed.
// All tables should have these following fields.
type Model struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"index" json:"created_at"`
	UpdatedAt time.Time `gorm:"index" json:"updated_at"`
}

type dbConfiguration struct {
	Host, User, Password, Database, Port string
}

func CreateDb(c *dbConfiguration) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Singapore",
		c.Host, c.User, c.Database, c.Port,
	)

	if c.Password != "" {
		dsn += fmt.Sprintf(" password=%s", c.Password)
	}

	var logMode logger.Interface
	// todo: config.IsDevelopment
	if true {
		logMode = logger.Default.LogMode(logger.Warn)
	} else {
		logMode = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logMode,
	})
	if err != nil {
		// error will be written automatically
		os.Exit(1)
	}

	return db
}

func NewDB() *gorm.DB {
	return CreateDb(&dbConfiguration{
		Host:     viper.GetString("db.hostname"),
		User:     viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Database: viper.GetString("db.database"),
		Port:     viper.GetString("db.port"),
	})
}
