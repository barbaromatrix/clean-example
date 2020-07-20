package presenter

import "github.com/barbaromatrix/clean-example/domain/model"

type userPresenter struct {}

type UserPresenter interface {
    ResponseUsers(us []*model.User) []*model.User
}

func NewUserPresenter() UserPresenter {
    return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Name = "Mr." + u.Name
	}
	return us
}
