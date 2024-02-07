package repositories

import (
	"errors"
	"simple-login/models"
	"sync"
)

type UserRepository interface {
	Save(user models.User) error
	ExistEmail(email string) bool
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
		return errors.New("유저 정보 등록 실패")
	}

	r.registerUsers[user.Email] = user

	return nil
}

func (r *memoryUserRepository) ExistEmail(email string) bool {

	_, existUser := r.registerUsers[email]

	return existUser
}
