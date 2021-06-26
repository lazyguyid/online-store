package core

import (
	"fmt"

	"github.com/labstack/echo"
)

// application struct
type application struct {
	usecases      []*ConfigUsecase
	repositories  map[string]Repository
	crepositories map[string]CRepository
	config        Config
	helper        interface{}
	storage       Storage
}

// Application func
func Application(cfg Config) App {
	app := new(application)
	app.config = cfg
	app.repositories = make(map[string]Repository)
	app.crepositories = make(map[string]CRepository)

	return app
}

func (app *application) Start() App {
	EchoStart(app)
	return app
}

func (app *application) Config() Config {
	return app.config
}

func (app *application) Storage() Storage {
	return app.storage
}

func (app *application) SetStorage(st Storage) {
	app.storage = st
}

func (app *application) Exit() {
	fmt.Println("shutdown")
}

func (app *application) Register(cu *ConfigUsecase) {
	app.usecases = append(app.usecases, cu)
}

func (app *application) Usecases() []*ConfigUsecase {
	return app.usecases
}

func (app *application) AddCRepository(name string, rp CRepository) {
	app.crepositories[name] = rp
}

func (app *application) AddRepository(name string, rp Repository) {
	app.repositories[name] = rp
}

func (app *application) Repositories() map[string]Repository {
	return app.repositories
}

func (app *application) CRepositories() map[string]CRepository {
	return app.crepositories
}

func (app *application) GetRepository(name string) Repository {
	return app.repositories[name]
}

func (app *application) GetCRepository(name string) CRepository {
	return app.crepositories[name]
}

func (app *application) RegisterHelper(h interface{}) {
	app.helper = h
}

func (app *application) Helper() interface{} {
	return app.helper
}

type ctx struct {
	echoCtx echo.Context
}

// NewContext func
func NewContext(cx ConfigContext) Context {
	c := new(ctx)
	c.echoCtx = cx.Echo
	return c
}

func (cx *ctx) Echo() echo.Context {
	return cx.echoCtx
}
