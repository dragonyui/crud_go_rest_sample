package models

// SKU (product's code), name, quantity dan price
type Product struct {
	Id int64 `json:"id"`
	Sku string `json:"sku"`
	Name string `json:"name"`
	Qty int `json:"qty"`
	Price float64 `json:"price"`
}
