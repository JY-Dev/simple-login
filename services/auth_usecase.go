package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"simple-login/dtos"
	"simple-login/repositories"
)

type AuthUseCase struct {
	userRepository repositories.UserRepository
}

func NewAuthUseCase() AuthUseCase {
	return AuthUseCase{
		userRepository: repositories.MemoryUserRepository(),
	}
}

func (u *AuthUseCase) LoginUser(loginUser dtos.LoginUserDto) error {

	user, err := u.userRepository.FindUser(loginUser.Email)
	if err != nil {
		return errors.New("유저가 존재하지 않습니다")
	}

	comparePassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if comparePassword != nil {
		return errors.New("유저 비밀번호가 일치하지 않습니다")
	}

	return nil
}
