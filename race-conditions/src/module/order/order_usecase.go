package order

import (
	"errors"
	"net/http"
	"online-store/core"
	"online-store/src/shared/domains"
	"online-store/src/shared/helpers"
)

type Usecase interface {
	Buy(c core.Context, data interface{}) (result core.Result)
}

type usecase struct {
	config      core.Config
	storage     core.Storage
	productRepo core.CRepository
	orderRepo   core.CRepository
	helper      helpers.Helper
}

func NewOrderUsecase(app core.App) Usecase {
	uc := new(usecase)
	uc.config = app.Config()
	uc.productRepo = app.GetCRepository("product")
	uc.orderRepo = app.GetCRepository("order")
	uc.storage = app.Storage()
	uc.helper = app.Helper().(helpers.Helper)

	return uc
}

func (uc *usecase) Buy(c core.Context, data interface{}) (result core.Result) {
	request := data.(*Request)

	order := new(domains.Order)
	order.UserID = request.UserID

	// transaction begin
	tx := uc.storage.Begin(core.StorageEngines.Postgres)

	defer func() {
		if r := recover(); r != nil {
			msg := uc.helper.IdentifyPanic("orderUsecase", r)
			result.Code = http.StatusBadRequest
			result.Message = "got panic"
			result.Error = errors.New(msg)
			tx.Rollback()
			return
		}
	}()

	if err := tx.Error; err != nil {
		return core.Result{
			Code:    http.StatusBadRequest,
			Message: "cannot do transaction",
			Error:   err,
		}
	}

	for _, p := range request.Products {
		// request stock item for order
		result = <-uc.productRepo.CustomFunc(&core.RepoParam{
			Fn:       "requestStockOrderWithTrx",
			UniqueID: p.ID,
			Data: map[string]int64{
				"qty": p.Qty,
			},
			Transaction: tx,
		})

		if result.Error != nil {
			result.Code = http.StatusBadRequest
			result.Message = "cannot request stock order"
			return result
		}

		product := result.Data.(*domains.Product)
		order.Products = append(order.Products, domains.OrderDetail{
			ProductID: p.ID,
			Price:     product.Price,
			Qty:       p.Qty,
		})
	}

	result = <-uc.orderRepo.Create(&core.RepoParam{
		Data:        order,
		Transaction: tx,
	})

	if result.Error != nil {
		result.Code = http.StatusBadRequest
		result.Message = "failed to create order"
		return result
	}

	if err := tx.Commit().Error; err != nil {
		return core.Result{
			Code:    http.StatusBadRequest,
			Message: "cannot commit transaction",
			Error:   err,
		}
	}

	result.Code = http.StatusOK
	result.Message = "success to create order"

	return result
}
