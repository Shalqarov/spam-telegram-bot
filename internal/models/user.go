package models

type User struct {
	Id        int64
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Service interface {
	CreateUser(user *User) (int64, error)
	GetUserById(id int64) (*User, error)
}

type Repo interface {
	CreateUser(user *User) (int64, error)
	GetUserById(id int64) (*User, error)
}
