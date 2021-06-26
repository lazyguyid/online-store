package repositories

import (
	"online-store/core"
	"online-store/src/shared/domains"
	"online-store/src/shared/helpers"

	"gorm.io/gorm"
)

type repository struct {
	conn   *gorm.DB
	helper helpers.Helper
	fn     map[string]func(params *core.RepoParam) <-chan core.Result
}

func NewOrderRepository(app core.App) core.CRepository {
	or := new(repository)
	or.conn = app.Storage().Postgres()
	return or
}

func (or *repository) CustomFunc(params *core.RepoParam) <-chan core.Result {
	return or.fn[params.Fn](params)
}

func (or *repository) Get(params *core.RepoParam) <-chan core.Result {
	output := make(chan core.Result, 0)
	go func() {
		defer close(output)
	}()

	return output
}

func (or *repository) Create(params *core.RepoParam) <-chan core.Result {
	if params.Transaction == nil {
		panic("cannot get transaction instance")
	}

	output := make(chan core.Result, 0)
	go func() {
		defer close(output)
		order := params.Data.(*domains.Order)
		err := params.Transaction.Create(&order).Error

		if err != nil {
			output <- core.Result{
				Error: err,
			}
			return
		}
		err = or.createOrderDetail(order, params)
		if err != nil {
			output <- core.Result{
				Error: err,
			}
			return
		}

		output <- core.Result{
			Data: order,
		}
		return
	}()

	return output
}

func (or *repository) Update(params *core.RepoParam) <-chan core.Result {
	output := make(chan core.Result, 0)
	go func() {
		defer close(output)
	}()

	return output
}

func (or *repository) Delete(params *core.RepoParam) <-chan core.Result {
	output := make(chan core.Result, 0)
	go func() {
		defer close(output)
	}()

	return output
}

func (or *repository) createOrderDetail(order *domains.Order, params *core.RepoParam) error {
	return params.Transaction.Create(&order.Products).Error
}
