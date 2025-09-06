package service

type UserService interface {
	GetAllUsers()
	CreateUsers()
	GetUserByUUID()
	UpdateUser()
	DeleteUser()
}
