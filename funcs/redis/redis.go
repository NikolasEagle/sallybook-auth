package redis

import (
	"context"
	"fmt"
	"os"
	"sallybook-auth/funcs/convert"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis"
)

var _ = godotenv.Load()

var (
	host string = os.Getenv("REDIS_HOST")

	port string = os.Getenv("REDIS_PORT")

	password string = os.Getenv("REDIS_PASSWORD")

	db int = convert.GetEnvAsInt(os.Getenv("REDIS_DB"))
)

var ctx = context.Background()

var client *redis.Client = redis.NewClient(&redis.Options{

	Addr: fmt.Sprintf("%s:%s", host, port),

	Password: password,

	DB: db,
})

func CheckConnection() {

}
