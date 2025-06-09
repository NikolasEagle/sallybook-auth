package redis_store

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sallybook-auth/funcs/convert"

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

func CheckConnection() error {

	_, err := Client.Ping(ctx).Result()

	if err != nil {

		slog.Error(err.Error())

		return fmt.Errorf("%s", "Error connecting to redis-store")

	}

	slog.Info("Successfully connected to redis-store!")

	return nil

}
