package model

// User Model
type User struct {
	Id   string
	Name string
}

// UserUsecase represent the User's usecases
type UserService interface {
	GetByID(id string) (User, error)
	CreateUser(user User) error
}

// UserRepository represent the User's repository contract
type UserRepository interface {
	GetByID(id string) (User, error)
	CreateUser(user User) error
}
