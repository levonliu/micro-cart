package model

type Cart struct {
	ID        int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	ProductID int64 `gorm:"not_null" json:"product_id"`
	UserID    int64 `gorm:"not_null" json:"user_id"`
	SizeID    int64 `gorm:"not_null" json:"size_id"`
	Num       int64 `gorm:"not_null" json:"num"`
}
