package structs

type User struct {
	FirstName string `form:"first_name"`

	SecondName string `form:"second_name"`

	Email string `form:"email"`

	Password string `form:"password"`
}
