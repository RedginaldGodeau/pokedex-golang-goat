package application

import (
	"backend/ent/gen"
	datastore "backend/pkg/datestore"
	"backend/pkg/env"
	"backend/pkg/router"
)

type Application struct {
	Router *router.Router
	DB     *gen.Client
	Env    *env.Env
}

func NewApplication() *Application {
	_env := env.NewEnv()

	return &Application{
		Router: router.NewRouter(_env.BackendPort),
		DB:     datastore.NewDatabase(_env),
		Env:    _env,
	}
}

func (app *Application) Start() {
	err := app.Router.Serve()
	if err != nil {
		return
	}
}
