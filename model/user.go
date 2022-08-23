package model

// User Model
type User struct {
	Id   string
	Name string
}

// UserUsecase represent the User's usecases
type UserUsecase interface {
	GetByID(id string) (Room, error)
}

// UserRepository represent the User's repository contract
type UserRepository interface {
	GetByID(id string) (Room, error)
}
