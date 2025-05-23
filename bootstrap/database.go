package bootstrap

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql" // GORM MySQL driver
	"gorm.io/gorm"         // GORM library
)

// NewMySQLDatabase creates a new GORM database connection
func NewMySQLDatabase(env *Env) *gorm.DB { // Changed return type to *gorm.DB
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName)
	
	log.Printf("Connecting to MySQL database with DSN: %s", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database with GORM: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting underlying sql.DB from GORM: %v", err)
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(env.DBMaxOpenConns)
	sqlDB.SetMaxIdleConns(env.DBMaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(env.DBConnMaxLifetime) * time.Minute)

	// GORM's Open function already pings the database, so explicit PingContext might be redundant
	// However, keeping a check here or relying on GORM's error handling during Open is fine.
	log.Println("Successfully connected to MySQL database using GORM.")
	return db
}

// CloseMySQLConnection closes the GORM database connection
func CloseMySQLConnection(db *gorm.DB) { // Changed parameter type to *gorm.DB
	if db == nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting underlying sql.DB from GORM for closing: %v", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Printf("Error closing MySQL database connection via GORM: %v", err)
	} else {
		log.Println("Connection to MySQL closed via GORM.")
	}
}
