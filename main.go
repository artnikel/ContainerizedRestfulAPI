package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type car struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Name  string `json:"name"`
	Age   int64  `json:"age"`
	Price int64  `json:"price"`
}

var cars = []car{
	{ID: "1", Brand: "Volkswagen", Name: "Jetta", Age: 2007, Price: 12000},
	{ID: "2", Brand: "BMW", Name: "X5", Age: 2009, Price: 20000},
	{ID: "3", Brand: "Mercedes-Benz", Name: "C63 AMG", Age: 2019, Price: 43000},
}

func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

func postCars(c *gin.Context) {
	var newCar car

	if err := c.BindJSON(&newCar); err != nil {
		return
	}

	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, newCar)
}

func getCarByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range cars {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}

func deleteCarByID(c *gin.Context) {
	id := c.Param("id")

	for i, a := range cars {
		if a.ID == id {
			cars = append(cars[:i], cars[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "car deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}

func updateCarByID(c *gin.Context) {
	id := c.Param("id")

	var updatedCar car
	for i := 0; i < len(cars); i++ {
		if cars[i].ID == id {
			if err := c.BindJSON(&updatedCar); err != nil {
				return
			}
			cars[i] = updatedCar
			c.IndentedJSON(http.StatusOK, updatedCar)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}

func main() {
	router := gin.Default()
	router.GET("/cars", getCars)
	router.GET("/cars/:id", getCarByID)
	router.POST("/cars", postCars)
	router.DELETE("/cars/:id", deleteCarByID)
	router.PUT("/cars/:id", updateCarByID)

	router.Run()
}
