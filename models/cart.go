package models

import "time"

type Cart struct {
	ID             int       `json:"id" gorm:"PRIMARY_KEY"`
	Product_ID     int       `json:"product_id"`
	Transaction_ID int       `json:"transaction_id"`
	Product        Product   `json:"product"`
	Topping        []Topping `gorm:"many2many:cart_toppings;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Topping_ID     []int     `json:"-" form:"topping_id" gorm:"-"`
	SubTotal       int       `json:"subtotal"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	UserID         int       `json:"user_id"`
}
