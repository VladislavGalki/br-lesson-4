package inmemory

import (
	userError "br-lesson-4/internal/domain/user/errors"
	userDomain "br-lesson-4/internal/domain/user/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (storage *Storage) GetUserList() ([]userDomain.User, error) {
	if len(storage.users) == 0 {
		return nil, userError.UserListEmptyError
	}

	return storage.users, nil
}

func (storage *Storage) GetUseByID(id string) (userDomain.User, error) {
	for _, user := range storage.users {
		if user.Id == id {
			return user, nil
		}
	}

	return userDomain.User{}, userError.UserNotFoundError
}

func (storage *Storage) GetUser(userReq userDomain.UserRequest) (userDomain.User, error) {
	for _, user := range storage.users {
		if user.Email == userReq.Email {
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password)); err != nil {
				return userDomain.User{}, userError.UserNotFoundError
			}

			return user, nil
		}
	}

	return userDomain.User{}, userError.UserNotFoundError
}

func (storage *Storage) CreateUser(domainUser userDomain.User) (userDomain.User, error) {
	for _, user := range storage.users {
		if user.Id == domainUser.Id {
			return userDomain.User{}, userError.UserAlreadyExistsError
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(domainUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return userDomain.User{}, err
	}

	domainUser.Id = uuid.NewString()
	domainUser.Password = string(hash)
	storage.users = append(storage.users, domainUser)
	return domainUser, nil
}

func (storage *Storage) UpdateUser(id string, domainUser userDomain.User) (userDomain.User, error) {
	for index, user := range storage.users {
		if user.Id == id {
			storage.users[index] = domainUser
			return domainUser, nil
		}
	}

	return userDomain.User{}, userError.UserNotFoundError
}

func (storage *Storage) DeleteUser(id string) error {
	for index, user := range storage.users {
		if user.Id == id {
			storage.users = append(storage.users[:index], storage.users[index+1:]...)
			return nil
		}
	}

	return userError.UserNotFoundError
}
