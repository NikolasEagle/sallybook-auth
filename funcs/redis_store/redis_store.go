package redis_store

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sallybook-auth/funcs/convert"

	redis_fiber "github.com/gofiber/storage/redis/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var _ = godotenv.Load()

var (
	host string = os.Getenv("REDIS_HOST")

	port string = os.Getenv("REDIS_PORT")

	password string = os.Getenv("REDIS_PASSWORD")

	db int = convert.GetEnvAsInt(os.Getenv("REDIS_DB"))
)

var ctx = context.Background()

var Client *redis.Client = redis.NewClient(&redis.Options{

	Addr: fmt.Sprintf("%s:%s", host, port),

	Password: password,

	DB: db,
})

var Store = redis_fiber.New(redis_fiber.Config{

	Host: host,

	Port: convert.GetEnvAsInt(port),

	Username: "",

	Password: password,

	Database: db,
})

func CheckConnection() error {

	_, err := Client.Ping(ctx).Result()

	if err != nil {

		slog.Error(err.Error())

		return fmt.Errorf("%s", "Error connecting to redis-store")

	}

	slog.Info("Successfully connected to redis-store!")

	return nil

}

func GetValue(key string) (string, error) {

	err := CheckConnection()

	if err != nil {

		return "", err

	}

	val, err := Client.Get(ctx, key).Result()

	if err == redis.Nil {

		return "none", nil

	} else if err != nil {

		slog.Error(err.Error())

		return "", fmt.Errorf("%s", "Error reading key of redis-store")

	} else {

		return val, nil

	}

}
