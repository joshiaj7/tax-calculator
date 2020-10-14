package controller

import (
	"github.com/joshiaj7/tax-calculator/internal/config"
	"github.com/joshiaj7/tax-calculator/internal/model"
	"github.com/joshiaj7/tax-calculator/internal/view"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// TypeMap will map item types
var TypeMap = map[int]string {
	1: "Food & Beverage",
	2: "Tobacco",
	3: "Entertainment",
}

// RefundableMap will map refundable conditions
var RefundableMap = map[int]bool {
	1: true,
	2: false,
	3: false,
}

// SetupRoute to handle routing
func SetupRoute(){
	fmt.Println("Handle Routes")
	http.HandleFunc("/", root)
	http.HandleFunc("/create", insertData)
	http.HandleFunc("/update", updateData)
	http.HandleFunc("/delete", deleteData)
	http.HandleFunc("/get", getData)
}

func root(w http.ResponseWriter, r *http.Request) {
	log.Println("This is root")
}

// insertData to insert data to db
func insertData(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var data model.Tax
    err := decoder.Decode(&data)
    if err != nil {
        panic(err)
    }

	// pass pointer of data to Create
	config.DB.Create(&data) 

	view.HTTPResponse(w, 200, "Success", data)
}

// updateData to update data from db
func updateData(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
    var item model.Tax
    err := decoder.Decode(&item)
    if err != nil {
        panic(err)
    }

	var existingItem model.Tax
	config.DB.Where("ID = ?", item.ID).First(&existingItem)

	existingItem.Name = item.Name
	existingItem.TaxCode = item.TaxCode
	existingItem.Price = item.Price
	config.DB.Save(&existingItem)

	view.HTTPResponse(w, 200, "Success", item)
}

// deleteData to delete data from db
func deleteData(w http.ResponseWriter, r *http.Request) {
	log.Println("Trying to get data")
	decoder := json.NewDecoder(r.Body)
    var item model.Tax
    err := decoder.Decode(&item)
    if err != nil {
        panic(err)
	}
	
	config.DB.Delete(&item, item.ID)

	view.HTTPResponse(w, 200, "Success", nil)
}

// getData to select data from db
func getData(w http.ResponseWriter, r *http.Request) {
	var data []model.Tax
	config.DB.Find(&data)
	
	for i := range data {
		// item type assignment
		data[i].Type = TypeMap[data[i].TaxCode]
		
		// item refundable condition
		data[i].Refundable = RefundableMap[data[i].TaxCode]

		// item tax calculation
		if data[i].TaxCode == 1 {
			data[i].Tax = data[i].Price * 0.1
		} else if data[i].TaxCode == 2 {
			data[i].Tax = 10.0 + (data[i].Price * 0.02)
		} else if data[i].TaxCode == 3 {
			if 0.0 < data[i].Price && data[i].Price < 100.0 {
				data[i].Tax = 0.0
			} else if data[i].Price >= 100.0 {
				data[i].Tax = 0.05 * (data[i].Price - 100.0)
			}
		}

		// item tax calculation
		data[i].Amount = data[i].Price + data[i].Tax
	}

	fmt.Println(data)

	view.HTTPResponse(w, 200, "Success", data)
}