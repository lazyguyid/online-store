package order

import "github.com/lazyguyid/gacor"

// Request struct
type Request struct {
	Products []Product `json:"products"`
	UserID   int64     `json:"userId"`
}

type Product struct {
	ID  int64 `json:"id"`
	Qty int64 `json:"qty"`
}

func NewOrderRequest() gacor.Request {
	return new(Request)
}

// GetInstance func
func (request *Request) GetInstance() interface{} {
	return request
}
