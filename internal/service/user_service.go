package service

import (
	"log"
	"user-management-api/internal/repository"
)

type userService struct {
	repo *repository.InMemoryUserRepository
}

func NewUserService(repo *repository.InMemoryUserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetAllUsers() {
	us.repo.FindAll()

	log.Println("GetAllUsers into userService")
}

func (us *userService) CreateUsers() {

}

func (us *userService) GetUserByUUID() {

}

func (us *userService) UpdateUser() {

}

func (us *userService) DeleteUser() {

}
