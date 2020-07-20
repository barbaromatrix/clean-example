package registry

import (
    "github.com/barbaromatrix/clean-example/interface/controller"
    ip "github.com/barbaromatrix/clean-example/interface/presenter"
    ir "github.com/barbaromatrix/clean-example/interface/repository"
    "github.com/barbaromatrix/clean-example/usecase/interactor"
    up "github.com/barbaromatrix/clean-example/usecase/presenter"
    ur "github.com/barbaromatrix/clean-example/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
