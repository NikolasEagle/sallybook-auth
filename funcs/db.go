package db

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

var host string = os.Getenv("POSTGRES_HOST")

var port string = os.Getenv("POSTGRES_PORT")

var user string = os.Getenv("POSTGRES_USER")

var password string = os.Getenv("POSTGRES_PASSWORD")

var dbname string = os.Getenv("POSTGRES_DB")

var psqlInfo string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func CheckConnection() {

	db, err := sql.Open("postgres", psqlInfo)

	fmt.Print(err)

	if err != nil {

		log.Fatal("Error opening connection to database")

		return
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {

		log.Fatal("Error ping to database")

		return
	}

	slog.Info("Successfully connected!")

}
