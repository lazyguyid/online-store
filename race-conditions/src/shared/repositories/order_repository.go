package repositories

import (
	"online-store/src/shared/domains"
	"online-store/src/shared/helpers"

	"github.com/lazyguyid/gacor"
	"gorm.io/gorm"
)

type repository struct {
	conn   *gorm.DB
	helper helpers.Helper
	fn     map[string]func(params *gacor.RepoParam) <-chan gacor.Result
}

func NewOrderRepository(app gacor.App) gacor.CRepository {
	or := new(repository)
	or.conn = app.Storage().Postgres()
	return or
}

func (or *repository) CustomFunc(params *gacor.RepoParam) <-chan gacor.Result {
	return or.fn[params.Fn](params)
}

func (or *repository) Get(params *gacor.RepoParam) <-chan gacor.Result {
	output := make(chan gacor.Result, 0)
	go func() {
		defer close(output)
	}()

	return output
}

func (or *repository) Create(params *gacor.RepoParam) <-chan gacor.Result {
	if params.Transaction == nil {
		panic("cannot get transaction instance")
	}

	output := make(chan gacor.Result, 0)
	go func() {
		defer close(output)
		order := params.Data.(*domains.Order)
		err := params.Transaction.Create(&order).Error

		if err != nil {
			output <- gacor.Result{
				Error: err,
			}
			return
		}
		err = or.createOrderDetail(order, params)
		if err != nil {
			output <- gacor.Result{
				Error: err,
			}
			return
		}

		output <- gacor.Result{
			Data: order,
		}
		return
	}()

	return output
}

func (or *repository) Update(params *gacor.RepoParam) <-chan gacor.Result {
	output := make(chan gacor.Result, 0)
	go func() {
		defer close(output)
	}()

	return output
}

func (or *repository) Delete(params *gacor.RepoParam) <-chan gacor.Result {
	output := make(chan gacor.Result, 0)
	go func() {
		defer close(output)
	}()

	return output
}

func (or *repository) createOrderDetail(order *domains.Order, params *gacor.RepoParam) error {
	return params.Transaction.Create(&order.Products).Error
}
