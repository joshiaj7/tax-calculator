package scripts

import (
	"fmt"

	"github.com/joshiaj7/tax-calculator/internal/config"
	"github.com/joshiaj7/tax-calculator/internal/model"
)

// Populate is to initially populate db
func Populate() {
	items := []model.Tax{
		{
			Name:    "Newspaper",
			TaxCode: 3,
			Price:   1000,
		},
		{
			Name:    "Nintendo 3DS",
			TaxCode: 3,
			Price:   300000,
		},
		{
			Name:    "Marlboro",
			TaxCode: 2,
			Price:   1500,
		},
		{
			Name:    "Gudang Garam",
			TaxCode: 2,
			Price:   1200,
		},
		{
			Name:    "Burger",
			TaxCode: 1,
			Price:   5500,
		},
		{
			Name:    "Coke",
			TaxCode: 1,
			Price:   1200,
		},
	}

	config.DB.Create(&items)

	fmt.Println("Database has been populated")
}
