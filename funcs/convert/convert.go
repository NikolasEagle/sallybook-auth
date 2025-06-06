package convert

import "strconv"

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

/*func GetUserStruct(data map[string]string) User {



}*/
