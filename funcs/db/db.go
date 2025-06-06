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

func OpenConnection() (*sql.DB, error) {

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {

		msg := "Error opening connection to database"

		slog.Error(msg)

		return nil, fmt.Errorf("%s", msg)
	}

	return db, nil

}

func CheckConnection() (*sql.DB, error) {

	db, err := OpenConnection()

	if err != nil {

		slog.Error(err.Error())

		return nil, err
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {

		msg := "Error ping to database"

		slog.Error(msg)

		return nil, fmt.Errorf("%s", msg)

	}

	return db, nil

}

func CheckPresenceUser(email string) error {

	_, err := CheckConnection()

	if err != nil {

		slog.Error(err.Error())

		return err

	}

	db, err := OpenConnection()

	if err != nil {

		slog.Error(err.Error())

		return err

	}

	defer db.Close()

	query := fmt.Sprintf(`SELECT first_name, second_name FROM users WHERE email='%s'`, email)

	rows, err := db.Query(query)

	if err != nil {

		msg := "Error selecting data from database"

		slog.Error(msg)

		return fmt.Errorf("%s", msg)

	}

	defer rows.Close()

	return nil

}
