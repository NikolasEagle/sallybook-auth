package convert

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
)

func GetEnvAsInt(value string) int {

	res, _ := strconv.Atoi(value)

	return res

}

type User struct {
	FirstName string

	SecondName string

	Email string

	Password string
}

func GetUserStruct(data map[string]string) User {

	var user User

	mapstructure.Decode(data, &user)

	return user

}
