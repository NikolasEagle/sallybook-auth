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

func CheckConnection() error {

	db, err := sql.Open("postgres", psqlInfo)

	var msg string

	if err != nil {

		msg = "Error opening connection to database"

		slog.Error(msg)

		return fmt.Errorf("%s", msg)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {

		msg = "Error ping to database"

		slog.Error(msg)

		return fmt.Errorf("%s", msg)

	}

	return nil

}
