package main

import (
	"github.com/joshiaj7/tax-calculator/internal/controller"
	"github.com/joshiaj7/tax-calculator/internal/config"

	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening on port 8080")
	
	err := config.SetupDB()

	fmt.Println(err)

	controller.SetupRoute()
	http.ListenAndServe(":8080", nil)
}