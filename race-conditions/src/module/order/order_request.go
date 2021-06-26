package order

import "online-store/core"

// Request struct
type Request struct {
	Products []Product `json:"products"`
	UserID   int64     `json:"userId"`
}

type Product struct {
	ID  int64 `json:"id"`
	Qty int64 `json:"qty"`
}

func NewOrderRequest() core.Request {
	return new(Request)
}

// GetInstance func
func (request *Request) GetInstance() interface{} {
	return request
}
