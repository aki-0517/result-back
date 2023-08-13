package main

import (
	"github.com/gin-gonic/gin"
)

type fruit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

var fruits = []fruit{
	{ID: 1, Name: "apple", Icon: "🍎"},
	{ID: 2, Name: "banana", Icon: "🍌"},
	{ID: 3, Name: "grape", Icon: "🍇"},
}

func main() {
	router := gin.Default()

	router.GET("/", getFruits)

	// サーバーを起動
	router.Run(":8080")
}

func getFruits(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(200, fruits)
}
