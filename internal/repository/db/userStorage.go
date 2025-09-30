package db

import (
	"br-lesson-4/internal/domain"
	userDomain "br-lesson-4/internal/domain/user/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type userStorage struct {
	db *pgx.Conn
}

func (s *userStorage) GetUserList() ([]userDomain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	var users []userDomain.User

	rows, err := s.db.Query(ctx, `SELECT * FROM users`)
	if err != nil {
		return []userDomain.User{}, nil
	}

	for rows.Next() {
		var user userDomain.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
			return []userDomain.User{}, nil
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return []userDomain.User{}, nil
	}

	return users, nil
}

func (s *userStorage) GetUseByID(id string) (userDomain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	var user userDomain.User
	row := s.db.QueryRow(ctx, `SELECT * FROM users WHERE id=$1`, id)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
		return userDomain.User{}, err
	}

	return user, nil
}

func (s *userStorage) GetUser(userReq userDomain.UserRequest) (userDomain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	var user userDomain.User
	row := s.db.QueryRow(ctx, `SELECT * FROM users WHERE email=$1`, userReq.Email)
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
		return userDomain.User{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password)); err != nil {
		return userDomain.User{}, err
	}
	return user, nil
}

func (s *userStorage) CreateUser(domainUser userDomain.User) (userDomain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	hash, err := bcrypt.GenerateFromPassword([]byte(domainUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return userDomain.User{}, err
	}

	domainUser.Id = uuid.NewString()
	domainUser.Password = string(hash)

	_, err = s.db.Exec(
		ctx,
		"INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)",
		domainUser.Id,
		domainUser.Name,
		domainUser.Email,
		domainUser.Password,
	)
	if err != nil {
		return userDomain.User{}, err
	}

	return domainUser, nil
}

func (s *userStorage) UpdateUser(id string, domainUser userDomain.User) (userDomain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	hash, err := bcrypt.GenerateFromPassword([]byte(domainUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return userDomain.User{}, err
	}

	domainUser.Id = id
	domainUser.Password = string(hash)

	_, err = s.db.Exec(ctx, "UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4",
		domainUser.Name,
		domainUser.Email,
		domainUser.Password,
		id,
	)
	if err != nil {
		return userDomain.User{}, err
	}

	return domainUser, nil
}

func (s *userStorage) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), domain.ContextTimeout)
	defer cancel()

	_, err := s.db.Exec(ctx, "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
