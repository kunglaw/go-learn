package main

import (
	"fmt"
	. "go-learn/simple-rest-api-concurrency/models"
	"log"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

var (
	products = []Product{
		{ID: 1, Name: "Product1", Category: 1, Price: 10000, Stock: 5},
		{ID: 2, Name: "Product2", Category: 2, Price: 15000, Stock: 6},
		{ID: 1, Name: "Product3", Category: 1, Price: 9000, Stock: 10},
	}
	mutex sync.Mutex
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{
			Message: "OK",
			Data:    "",
		})
	})
	e.GET("products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)
	})
	e.GET("products/:id", getProduct)

	log.Fatal(e.Start(":5678"))
}

func getProduct(c echo.Context) error {
	id := c.Param("id")

	for _, product := range products {
		if fmt.Sprint(product.ID) == id {
			return c.JSON(http.StatusOK, product)
		}
	}

	return echo.ErrNotFound
}

func updateProduct() {

}
