package repositories

import (
	"errors"
	"fmt"
	"simple-login/models"
	"sync"
)

type UserRepository interface {
	Save(user models.User) error
	ExistEmail(email string) bool
	FindUser(email string) (models.User, error)
}

type memoryUserRepository struct {
	registerUsers map[string]models.User
	nextID        int64
	mu            sync.Mutex
}

var (
	once           sync.Once
	userRepository UserRepository
)

func MemoryUserRepository() UserRepository {
	once.Do(func() {
		userRepository = &memoryUserRepository{
			registerUsers: make(map[string]models.User),
		}
	})
	return userRepository
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

func (r *memoryUserRepository) FindUser(email string) (models.User, error) {
	fmt.Println(len(r.registerUsers))
	user, exist := r.registerUsers[email]
	if exist {
		return user, nil
	}

	return models.User{}, errors.New("유저가 존재하지 않습니다")
}
