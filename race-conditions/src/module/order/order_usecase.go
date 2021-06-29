package order

import (
	"errors"
	"net/http"
	"online-store/src/shared/domains"
	"online-store/src/shared/helpers"

	"github.com/lazyguyid/gacor"
)

type Usecase interface {
	Buy(c gacor.Context, data interface{}) (result gacor.Result)
}

type usecase struct {
	config      gacor.Config
	storage     gacor.Storage
	productRepo gacor.CRepository
	orderRepo   gacor.CRepository
	helper      helpers.Helper
}

func NewOrderUsecase(app gacor.App) Usecase {
	uc := new(usecase)
	uc.config = app.Config()
	uc.productRepo = app.GetCRepository("product")
	uc.orderRepo = app.GetCRepository("order")
	uc.storage = app.Storage()
	uc.helper = app.Helper().(helpers.Helper)

	return uc
}

func (uc *usecase) Buy(c gacor.Context, data interface{}) (result gacor.Result) {
	request := data.(*Request)

	order := new(domains.Order)
	order.UserID = request.UserID

	// transaction begin
	tx := uc.storage.Begin(gacor.StorageEngines.Postgres)

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
		return gacor.Result{
			Code:    http.StatusBadRequest,
			Message: "cannot do transaction",
			Error:   err,
		}
	}

	for _, p := range request.Products {
		// request stock item for order
		result = <-uc.productRepo.CustomFunc(&gacor.RepoParam{
			Fn:       "requestStockOrderWithTrx",
			UniqueID: p.ID,
			Data: map[string]int64{
				"qty": p.Qty,
			},
			Transaction: tx,
		})

		if result.Error != nil {
			tx.Rollback()
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

	result = <-uc.orderRepo.Create(&gacor.RepoParam{
		Data:        order,
		Transaction: tx,
	})

	if result.Error != nil {
		tx.Rollback()
		result.Code = http.StatusBadRequest
		result.Message = "failed to create order"
		return result
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return gacor.Result{
			Code:    http.StatusBadRequest,
			Message: "cannot commit transaction",
			Error:   err,
		}
	}

	result.Code = http.StatusOK
	result.Message = "success to create order"

	return result
}
