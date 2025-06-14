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

		slog.Error(err.Error())

		return nil, fmt.Errorf("%s", "Error opening connection to database")
	}

	return db, nil

}

func CheckConnection() (*sql.DB, error) {

	db, err := OpenConnection()

	if err != nil {

		return nil, err
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {

		slog.Error(err.Error())

		return nil, fmt.Errorf("%s", "Error ping to database")

	}

	return db, nil

}

func CheckPresenceUser(email string) (bool, error) {

	_, err := CheckConnection()

	if err != nil {

		return false, err

	}

	db, err := OpenConnection()

	if err != nil {

		return false, err

	}

	defer db.Close()

	query := fmt.Sprintf(`SELECT id FROM users WHERE email='%s'`, email)

	rows, err := db.Query(query)

	if err != nil {

		slog.Error(err.Error())

		return false, fmt.Errorf("%s", "Error selecting data from database")

	}

	defer rows.Close()

	users := []structs.User{}

	for rows.Next() {

		user := structs.User{}

		err := rows.Scan(&user.Id)

		if err != nil {

			msg := "Error scanning row from database"

			slog.Error(msg)

			continue

		}

		users = append(users, user)

	}

	return len(users) > 0, nil

}

func CreateUser(firstName, secondName, email, password string) (string, error) {

	_, err := CheckConnection()

	if err != nil {

		return "", err

	}

	db, err := OpenConnection()

	if err != nil {

		return "", err

	}

	defer db.Close()

	id := uuid.GetUUID()

	hash, _ := pw.HashPassword(password)

	query := fmt.Sprintf(`INSERT INTO users VALUES ('%s', '%s', '%s', '%s', '%s')`, id, firstName, secondName, email, hash)

	rows, err := db.Query(query)

	if err != nil {

		slog.Error(err.Error())

		return "", fmt.Errorf("%s", "Error creating data into database")

	}

	defer rows.Close()

	return email, nil

}

func CheckPassword(email, password string) (bool, error) {

	_, err := CheckConnection()

	if err != nil {

		return false, err

	}

	db, err := OpenConnection()

	if err != nil {

		return false, err

	}

	defer db.Close()

	query := fmt.Sprintf("SELECT hash FROM users WHERE email='%s'", email)

	rows, err := db.Query(query)

	if err != nil {

		slog.Error(err.Error())

		return false, fmt.Errorf("%s", "Error selecting data from database")

	}

	defer rows.Close()

	users := []structs.User{}

	for rows.Next() {

		user := structs.User{}

		err := rows.Scan(&user.Hash)

		if err != nil {

			slog.Error(err.Error())

			continue

		}

		users = append(users, user)

	}

	correctPassword := pw.CheckPasswordHash(password, users[0].Hash)

	return correctPassword, nil

}

func GetUserInfo(email string) (*structs.User, error) {
	return nil, nil
}
