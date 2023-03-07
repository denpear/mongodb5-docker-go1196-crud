package common

import (
	_ "encoding/json"
	"fmt"
	"log"
	"mongo-docker-go-crud/bookmarkapi/model"
	_ "net/http"
	"os"

	_ "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	PostgresConn *gorm.DB
	err          error
)

func initPostrgesDB() {
	var (
		host     = getEnvVariable("DB_HOST")
		port     = getEnvVariable("DB_PORT")
		user     = getEnvVariable("DB_USER")
		dbname   = getEnvVariable("DB_NAME")
		password = getEnvVariable("DB_PASSWORD")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	PostgresConn, err = gorm.Open("postgres", conn)
	PostgresConn.AutoMigrate(model.BookmarkPdb{})

	if err != nil {
		log.Fatal(err)
	}
}

func getEnvVariable(key string) string {
	err := godotenv.Load("bookmarkapi/common/postgresDB.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	return os.Getenv(key)
}
