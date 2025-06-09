package structs

type User struct {
	Id string

	FirstName string `form:"first_name"`

	SecondName string `form:"second_name"`

	Email string `form:"email"`

	Password string `form:"password"`

	Hash string
}

type SessionStore struct {
	Email string
}
