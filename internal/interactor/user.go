package interactor

import (
	"github.com/xxlifestyle/auth_service/internal/entity"
	"github.com/xxlifestyle/auth_service/internal/repository"
	"github.com/xxlifestyle/auth_service/internal/utils"
)

type userInteractor struct {
	userRepository repository.UserRepository
}

type UserInteractor interface {
	CreateUser(entity.CreateUserDTO) (*entity.User, error)
	GetUser(id string) (*entity.User, error)
}

func (ui userInteractor) CreateUser(u entity.CreateUserDTO) (*entity.User, error) {
	hashPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	newUser := entity.User{
		Username:         u.Username,
		PasswordHash:     hashPass,
		BirthdayDate:     u.BirthdayDate,
		Email:            u.Email,
		IsEmailConfirmed: false,
		Role:             "USER",
	}
	err = ui.userRepository.Create(&newUser)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}
func (ui userInteractor) GetUser(id string) (*entity.User, error) {
	user, err := ui.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserInteractor(userRepository repository.UserRepository) UserInteractor {
	return &userInteractor{
		userRepository: userRepository,
	}
}
