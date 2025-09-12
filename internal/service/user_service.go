package service

import (
	"user-management-api/internal/models"
	"user-management-api/internal/repository"
	"user-management-api/internal/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (us *userService) GetAllUsers(search string, page, limit int) ([]models.User, error) {
	users, err := us.repo.FindAll()
	if err != nil {
		return nil, utils.WrapError("failed to fetch users", utils.ErrCodeInternal, err)
	}

	return users, nil
}

func (us *userService) CreateUsers(user models.User) (models.User, error) {
	user.Email = utils.NormalizeString(user.Email)

	if _, exists := us.repo.FindByEmail(user.Email); exists {
		return models.User{}, utils.NewError("email already exitst", utils.ErrCodeConflict)
	}

	user.UUID = uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, utils.WrapError("failed to hash password", utils.ErrCodeInternal, err)
	}

	user.Password = string(hashedPassword)

	if err := us.repo.Create(user); err != nil {
		return models.User{}, utils.WrapError("failed to create password", utils.ErrCodeInternal, err)
	}

	return user, nil

}

func (us *userService) GetUserByUUID(uuid string) (models.User, error) {
	user, found := us.repo.FindByUUID(uuid)
	if !found {
		return models.User{}, utils.NewError("user not found", utils.ErrCodeNotFound)
	}

	return user, nil
}

func (us *userService) UpdateUser() {

}

func (us *userService) DeleteUser() {

}
