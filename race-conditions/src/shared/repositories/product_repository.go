package repositories

import (
	"fmt"
	"online-store/core"
	"online-store/src/shared/domains"
	"online-store/src/shared/helpers"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type productRepo struct {
	conn   *gorm.DB
	helper helpers.Helper
	fn     map[string]func(params *core.RepoParam) <-chan core.Result
}

func NewProductRepository(app core.App) core.CRepository {
	pr := new(productRepo)
	pr.conn = app.Storage().Postgres()

	pr.fn = make(map[string]func(params *core.RepoParam) <-chan core.Result)
	pr.fn["requestStockOrderWithTrx"] = pr.requestStockOrderWithTrx

	return pr
}

func (pr *productRepo) Create(rp *core.RepoParam) <-chan core.Result {
	result := make(chan core.Result)

	go func() {
		defer close(result)
	}()

	return result
}

func (pr *productRepo) Get(rp *core.RepoParam) <-chan core.Result {
	result := make(chan core.Result)

	go func() {
		defer close(result)

	}()

	return result
}

func (pr *productRepo) Update(rp *core.RepoParam) <-chan core.Result {
	result := make(chan core.Result)

	go func() {
		defer close(result)
	}()

	return result
}

func (pr *productRepo) Delete(rp *core.RepoParam) <-chan core.Result {
	result := make(chan core.Result)

	go func() {
		defer close(result)
	}()

	return result
}

func (pr *productRepo) CustomFunc(params *core.RepoParam) <-chan core.Result {
	return pr.fn[params.Fn](params)
}

func (pr *productRepo) requestStockOrderWithTrx(params *core.RepoParam) <-chan core.Result {
	if params.Transaction == nil {
		panic("cannot get transaction instance")
	}

	output := make(chan core.Result, 0)
	go func() {
		defer close(output)
		product := new(domains.Product)
		data := params.Data.(map[string]int64)
		// lock the row with specific condition
		err := params.Transaction.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id", params.UniqueID).Where(fmt.Sprintf("qty >= %d", data["qty"])).First(&product).Error
		if err != nil {
			output <- core.Result{
				Error: err,
			}
			return
		}
		// minus qty product
		product.Qty -= data["qty"]

		err = params.Transaction.Table("products").Where("id", product.ID).Updates(map[string]interface{}{
			"qty": product.Qty,
		}).Error

		if err != nil {
			output <- core.Result{
				Error: err,
			}
			return
		}

		output <- core.Result{
			Data: product,
		}

		return
	}()

	return output

}
