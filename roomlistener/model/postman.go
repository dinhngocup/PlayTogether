package model

// ConnectionManager store map <userId, []connectionId>
type Postman struct {
	UserConnectionMap map[string][]string
}

func NewPostman() *Postman {
	return &Postman{
		UserConnectionMap: make(map[string][]string),
	}
}

type PostmanService interface {
	MapUserConnection(userId string, connectionId string)
	DeliverMessage(userId string, message string)
}
