package service

import "user-management-api/internal/models"

type UserService interface {
	GetAllUsers()
	CreateUsers(user models.User)
	GetUserByUUID()
	UpdateUser()
	DeleteUser()
}
