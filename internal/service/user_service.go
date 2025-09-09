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

func (us *userService) CreateUsers(user models.User) (models.User, error) {
	user.Email = utils.NormalizeString(user.Email)

	if _, exists := us.repo.FindByEmail(user.Email); exists {
		return models.User{}, utils.NewError("email already exitst", utils.ErrCodeConflict)
	}
}

func (us *userService) GetUserByUUID() {

}

func (us *userService) UpdateUser() {

}

func (us *userService) DeleteUser() {

}
