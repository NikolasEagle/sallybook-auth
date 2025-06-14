package structs

type User struct {
	Id string `json:"-"`

	FirstName string `form:"first_name" json:"first_name"`

	SecondName string `form:"second_name" json:"second_name"`

	Email string `form:"email" json:"email"`

	Password string `form:"password" json:"-"`

	Hash string `json:"-"`
}

type SessionStore struct {
	Email string
}
