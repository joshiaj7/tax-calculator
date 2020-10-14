package model

// Tax is table model
type Tax struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `json:"name"`
	TaxCode int `json:"tax_code"`
	Type string `gorm:"-" json:"type"` 
	Refundable bool `gorm:"-" json:"refundable"` 
	Price float64 `json:"price"`
	Tax float64 `gorm:"-" json:"tax"` 
	Amount float64 `gorm:"-" json:"amount"` 
}