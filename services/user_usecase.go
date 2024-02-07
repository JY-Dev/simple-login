package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"simple-login/dtos"
	"simple-login/models"
	"simple-login/repositories"
)

type UserUseCase struct {
	userRepository repositories.UserRepository
}

func NewUserUseCase() UserUseCase {
	return UserUseCase{
		userRepository: repositories.MemoryUserRepository(),
	}
}

func (u *UserUseCase) CreateUser(userDto dtos.CreateUserDto) error {

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)

	newUser := &models.User{
		Nickname: userDto.Nickname,
		Email:    userDto.Email,
		Password: string(encryptedPassword),
	}

	err := u.userRepository.Save(*newUser)
	if err != nil {
		return errors.New("이미 등록된 유저이기 때문에 등록 할 수 없습니다.")
	}

	return nil
}
