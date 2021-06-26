package domains

import "time"

// Product struct
type Product struct {
	ID      int64      `json:"id" gorm:"id"`
	Name    string     `json:"name" gorm:"name"`
	Price   float64    `json:"price" gorm:"price"`
	Qty     int64      `json:"qty" gorm:"qty"`
	Created time.Time  `json:"created" gorm:"created"`
	Updated *time.Time `json:"updated" gorm:"updated"`
}
