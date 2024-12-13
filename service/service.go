package service

import (
	"context"
	"testingExample/database"
)

type UserModel struct {
	Name string
}

type Service interface {
	GetAll(ctx context.Context) ([]UserModel, error)
}

type service struct {
	db database.Database
}

func NewService(db database.Database) Service {
	return &service{db: db}
}

func (s *service) GetAll(ctx context.Context) ([]UserModel, error) {

	users := make([]UserModel, 0)

	conn := s.db.GetConnection()

	if err := conn.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
