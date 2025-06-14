package structs

type User struct {
	Id string

	FirstName string `form:"first_name" json:"first_name"`

	SecondName string `form:"second_name" json:"second_name"`

	Email string `form:"email" json:"email"`

	Password string `form:"password"`

	Hash string
}

type SessionStore struct {
	Email string
}
