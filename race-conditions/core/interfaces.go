package core

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (

	// Context interface
	Context interface {
		Echo() echo.Context
	}

	// Config interface
	Config interface {
		Set(k string, v interface{})
		Get(k string) interface{}
	}

	// Storage interface
	Storage interface {
		Postgres() (db *gorm.DB)
		Begin(storageEngine string) (db *gorm.DB)
	}

	// App interface
	App interface {
		Start() App
		Config() Config
		Exit()
		Storage() Storage
		SetStorage(Storage)
		Register(cu *ConfigUsecase)
		Usecases() []*ConfigUsecase
		AddCRepository(name string, rp CRepository)
		AddRepository(name string, rp Repository)
		GetCRepository(name string) CRepository
		GetRepository(name string) Repository
		Repositories() map[string]Repository
		CRepositories() map[string]CRepository
		RegisterHelper(h interface{})
		Helper() interface{}
	}

	// Rest interface
	Rest interface {
		Mount(routes interface{})
		Path() string
	}

	// API interface
	API interface {
		Connect() error
		WhenDown(func()) error
		WhenUp(func()) error
		Post(p *APIParam) Result
		Put(p *APIParam) Result
		Get(p *APIParam) Result
		Delete(p *APIParam) Result
		Option(p *APIParam) Result
	}

	// CAPI interface
	CAPI interface {
		Connect() error
		WhenDown(func()) error
		WhenUp(func()) error
		Post(p *APIParam) <-chan Result
		Put(p *APIParam) <-chan Result
		Get(p *APIParam) <-chan Result
		Delete(p *APIParam) <-chan Result
		Option(p *APIParam) <-chan Result
	}

	// Repository interface
	Repository interface {
		Get(rp *RepoParam) Result
		Create(rp *RepoParam) Result
		Update(rp *RepoParam) Result
		Delete(rp *RepoParam) Result
		CustomFunc(rp *RepoParam) Result
	}

	// CRepository interface
	CRepository interface {
		Get(rp *RepoParam) <-chan Result
		Create(rp *RepoParam) <-chan Result
		Update(rp *RepoParam) <-chan Result
		Delete(rp *RepoParam) <-chan Result
		CustomFunc(rp *RepoParam) <-chan Result
	}

	// Usecase interface
	Usecase interface {
		Exec(c Context, d interface{}) Result
	}

	// Request interface
	Request interface {
		EchoBinder(c echo.Context) (err error)
		GetInstance() interface{}
	}

	// HTTPServer interface
	HTTPServer interface {
		Start() App
		Shutdown() error
	}
)
