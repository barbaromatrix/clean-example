package interactor

import (
	"github.com/barbaromatrix/clean-example/domain/model"
	"github.com/barbaromatrix/clean-example/usecase/presenter"
	"github.com/barbaromatrix/clean-example/usecase/repository"
)

// this is responsible for allowing dependency injection
type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	GET(u []*model.User) ([]*model.User, error)
}

// it receives an {repository, presenter} and returns an instance of userInteractor (responsible for input port)
func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (us *userInteractor) GET(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}
