package domain

type UserModel struct {
	ID      uint64
	Name    string
	Email   string
	Message string
}

func NewUserModel(name, email string) *UserModel {
	return &UserModel{Name: name, Email: email, Message: "hello"}
}
