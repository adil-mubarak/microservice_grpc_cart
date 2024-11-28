package models

type Cart struct {
	ID        uint `gorm:"primarykey"`
	UserID    int  `gorm:"not null;int"`
	ProductID int  `gorm:"not null; int"`
	Quantity  int  `gorm:"not null; int"`
}
