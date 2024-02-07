package repositories

import (
	"errors"
	"simple-login/models"
	"sync"
)

type UserRepository interface {
	Save(user models.User) error
	ExistEmail(email string) bool
	FindUser(email string) models.User
}

type memoryUserRepository struct {
	registerUsers map[string]models.User
	nextID        int64
	mu            sync.Mutex
}

func NewMemoryUserRepository() UserRepository {
	return &memoryUserRepository{
		registerUsers: make(map[string]models.User),
	}
}

func (r *memoryUserRepository) Save(user models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.ExistEmail(user.Email) {
		return errors.New("해당 이메일로 등록된 유저가 존재합니다")
	}

	r.registerUsers[user.Email] = user

	return nil
}

func (r *memoryUserRepository) ExistEmail(email string) bool {

	_, existUser := r.registerUsers[email]

	return existUser
}

func (r *memoryUserRepository) FindUser(email string) models.User {

	user, _ := r.registerUsers[email]

	return user
}
