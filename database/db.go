package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

//edit to export env var
func ConnectWithDB() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	//
	DB_HOST := os.Getenv("DATABASE_URL")
	// DB_PASS := os.Getenv("DATABASE_PASS")
	// DB_URL := os.Getenv("DATABASE_HOST")
	// DB_USER := os.Getenv("DATABASE_USER")
	// DB_DB := os.Getenv("DATABASE_DB")
	// connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", DB_URL, DB_USER, DB_PASS, DB_DB)
	DB, err = gorm.Open(postgres.Open(DB_HOST))
	if err != nil {
		log.Panic("ERROR - Connection with db fail")
	}

	fmt.Println("Connected")
}
