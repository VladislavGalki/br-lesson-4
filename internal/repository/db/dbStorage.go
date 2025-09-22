package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"
	"log"
)

type Storage struct {
	taskStorage
	userStorage
}

func NewStorage(connectionPath string) (*Storage, error) {
	db, err := pgx.Connect(context.Background(), connectionPath)
	if err != nil {
		return nil, err
	}

	return &Storage{
		taskStorage{db},
		userStorage{db},
	}, nil
}

func Migrations(dsn string, migratePath string) error {
	mPath := fmt.Sprintf("file://%s", migratePath)
	m, err := migrate.New(mPath, dsn)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
		log.Println("Migrations complete")
	}

	return nil
}
