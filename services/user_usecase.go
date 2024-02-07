package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"simple-login/dtos"
	"simple-login/models"
	"simple-login/repositories"
)

type UserUseCase struct {
	userRepository repositories.UserRepository
}

func NewUserUseCase() UserUseCase {
	return UserUseCase{
		userRepository: repositories.NewMemoryUserRepository(),
	}
}

func (u *UserUseCase) CreateUser(userDto dtos.CreateUserDto) error {

	encryptedPassword := encrypted(userDto.Password)

	newUser := &models.User{
		Nickname: userDto.Nickname,
		Email:    userDto.Email,
		Password: encryptedPassword,
	}

	err := u.userRepository.Save(*newUser)
	if err != nil {
		return errors.New("이미 등록된 유저이기 때문에 등록 할 수 없습니다.")
	}

	return nil
}

func encrypted(password string) string {

	passwordBytes := []byte(password)

	hash := sha256.Sum256(passwordBytes)

	return hex.EncodeToString(hash[:])
}
