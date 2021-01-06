package main

import (
	"github.com/joshiaj7/tax-calculator/internal/config"
	"github.com/joshiaj7/tax-calculator/internal/controller"
	"github.com/joshiaj7/tax-calculator/internal/scripts"

	"fmt"
	"net/http"
)

func main() {
	err := config.SetupDB()
	fmt.Println(err)

	scripts.Populate()

	controller.SetupRoute()
	http.ListenAndServe(":8080", nil)

	fmt.Println("Listening on port 8080")
}
