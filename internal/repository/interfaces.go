package repository

import "user-management-api/internal/models"

type UserRepository interface {
	FindAll()
	Create()
	FindByUUID()
	Update()
	Delete()
	FindByEmail(email string) (models.User, bool)
}
