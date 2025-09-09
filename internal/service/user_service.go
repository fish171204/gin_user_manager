package service

import (
	"user-management-api/internal/models"
	"user-management-api/internal/repository"
	"user-management-api/internal/utils"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetAllUsers() {

}

func (us *userService) CreateUsers(user models.User) {
	user.Email = utils.NormalizeString(user.Email)
}

func (us *userService) GetUserByUUID() {

}

func (us *userService) UpdateUser() {

}

func (us *userService) DeleteUser() {

}
