package repository

import "github.com/barbaromatrix/clean-example/domain/model"

type UserRepository interface {
    FindAll(u []*model.User) ([]*model.User, error)
}

