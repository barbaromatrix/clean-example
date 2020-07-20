package presenter

import "github.com/barbaromatrix/clean-example/domain/model"

type UserPresenter interface {
    ResponseUsers(u []*model.User) []*model.User
}
