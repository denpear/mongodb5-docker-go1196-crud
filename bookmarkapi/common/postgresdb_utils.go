package common

import (
	_ "encoding/json"
	"fmt"
	_ "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"mongo-docker-go-crud/bookmarkapi/model"
	_ "net/http"
	"os"
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
	PostgresConn.AutoMigrate(model.Applicant{}, model.AdditionalOptions{})
	//PostgresConn.Model(&model.Applicant{}).AddForeignKey("id", "additional_options(id)", "CASCADE", "CASCADE") - n/a
	//PostgresConn.Model(&model.AdditionalOption{}).AddForeignKey("id", "additional_options(id)", "CASCADE", "CASCADE")
	//PostgresConn.Model(&model.AdditionalOption{}).AddForeignKey("id", "additional_options(applicant_id)", "CASCADE", "CASCADE")

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
