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

	PASS := os.Getenv("PASS")
	DB_URL := os.Getenv("DATABASE_URL")
	connectionString := fmt.Sprintf("host=%s user=root password=%s dbname=root port=5432 sslmode=disable", DB_URL, PASS)
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("ERROR - Connection with db fail")
	}
	fmt.Println("Connected")
}
