package main

import (
	"api_assignment/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handlers.RegisterRoutes(r)
	r.Run(":8080")
}
