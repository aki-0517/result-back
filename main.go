package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type player struct {
	Name    string `json:"name"`
	Point   int    `json:"point"`
	Sail_no int    `json:"sail_no"`
	Sex     string `json:"sex"`
}

var players = []player{
	{Name: "ユーザー11", Point: 100, Sail_no: 17, Sex: "man"},
	{Name: "ユーザー2", Point: 90, Sail_no: 17, Sex: "woman"},
	{Name: "ユーザー3", Point: 85, Sail_no: 17, Sex: "man"},
}

func main() {
	router := gin.Default()

	router.GET("/", getPlayers)

	public := router.Group("/api")

    public.POST("/register", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "data": "this is the register endpoint.",
        })
    })

	// サーバーを起動
	router.Run(":8000")
}

func getPlayers(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(200, players)
}
