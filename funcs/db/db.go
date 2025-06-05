package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"sallybook-auth/funcs/convert"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var _ = godotenv.Load()

var (
	host string = os.Getenv("POSTGRES_HOST")

	port int = convert.GetEnvAsInt(os.Getenv("POSTGRES_PORT"))

	user string = os.Getenv("POSTGRES_USER")

	password string = os.Getenv("POSTGRES_PASSWORD")

	dbname string = os.Getenv("POSTGRES_DB")
)

var psqlInfo string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func CheckConnection() {

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {

		slog.Error("Error opening connection to database")

		return
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {

		slog.Error("Error ping to database")

		return
	}

	slog.Info("Successfully connected to database!")

}
