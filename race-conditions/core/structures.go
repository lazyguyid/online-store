package core

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type (
	// Result struct
	Result struct {
		Data    interface{} `json:"data"`
		Error   error       `json:"error"`
		Message string      `json:"message"`
		Code    int         `json:"code"`
	}

	// APIParam struct
	APIParam struct {
		Headers map[string]string
		Data    map[string]interface{}
		Path    string
		Binder  interface{}
		URL     *string
	}

	// RepoParam struct
	RepoParam struct {
		Fn          string
		Data        interface{}
		UniqueID    interface{}
		Context     Context
		Binder      interface{}
		Transaction *gorm.DB
	}

	// Middleware struct
	Middleware struct {
		Echo func(c echo.HandlerFunc) echo.HandlerFunc
	}

	// ConfigUsecase struct
	ConfigUsecase struct {
		RestPath      string
		Topics        []string
		Usecase       func(c Context, request interface{}) Result
		RequestParser func() Request
		Version       string
		Enable        bool
		HTTPMethod    string
		Middleware    *Middleware
	}

	// ConfigContext struct
	ConfigContext struct {
		Echo echo.Context
	}
)
