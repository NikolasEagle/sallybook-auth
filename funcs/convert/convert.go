package convert

import (
	"strconv"
)

func GetEnvAsInt(value string) int {

	res, _ := strconv.Atoi(value)

	return res

}
