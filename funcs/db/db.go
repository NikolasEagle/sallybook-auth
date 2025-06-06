package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"sallybook-auth/funcs/convert"
	"sallybook-auth/funcs/pw"
	"sallybook-auth/funcs/uuid"
	"sallybook-auth/structs"

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

func CheckPresenceUser(email string) (bool, error) {

	_, err := CheckConnection()

	if err != nil {

		slog.Error(err.Error())

		return false, err

	}

	db, err := OpenConnection()

	if err != nil {

		slog.Error(err.Error())

		return false, err

	}

	defer db.Close()

	query := fmt.Sprintf(`SELECT first_name, second_name FROM users WHERE email='%s'`, email)

	rows, err := db.Query(query)

	if err != nil {

		msg := "Error selecting data from database"

		slog.Error(msg)

		return false, fmt.Errorf("%s", msg)

	}

	defer rows.Close()

	users := []structs.User{}

	for rows.Next() {

		user := structs.User{}

		err := rows.Scan(&user.FirstName, &user.SecondName)

		if err != nil {

			msg := "Error scanning row from database"

			slog.Error(msg)

			continue

		}

		users = append(users, user)

	}

	if len(users) > 0 {

		return true, nil

	}

	return false, nil

}

func CreateUser(firstName, secondName, email, password string) (string, error) {

	_, err := CheckConnection()

	if err != nil {

		slog.Error(err.Error())

		return "", err

	}

	db, err := OpenConnection()

	if err != nil {

		slog.Error(err.Error())

		return "", err

	}

	defer db.Close()

	id := uuid.GetUUID()

	hash, _ := pw.HashPassword(password)

	query := fmt.Sprintf(`INSERT INTO users VALUES ('%s', '%s', '%s', '%s', '%s')`, id, firstName, secondName, email, hash)

	rows, err := db.Query(query)

	if err != nil {

		msg := "Error creating data into database"

		slog.Error(msg)

		return "", fmt.Errorf("%s", msg)

	}

	defer rows.Close()

	return email, nil

}
