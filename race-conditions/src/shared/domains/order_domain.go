package domains

// Order struct
type Order struct {
	ID       int64         `json:"id" gorm:"id"`
	Products []OrderDetail `json:"products" gorm:"-"`
	UserID   int64         `json:"userId" json:"userId"`
	Total    int64         `json:"total"`
}

// OrderDetail struct
type OrderDetail struct {
	ID        int64   `json:"id"`
	ProductID int64   `json:"productId"`
	Qty       int64   `json:"qty"`
	Price     float64 `json:"price"`
}
