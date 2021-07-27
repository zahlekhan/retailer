package models

type OrderBook struct {
	OrderID   uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Quantity  uint
}

func (o *OrderBook) TableName() string {
	return "order_book"
}
