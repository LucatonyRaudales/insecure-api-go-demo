package main

import (
	"github.com/gin-gonic/gin"
	"insecure-api-go/user"
)

func main() {
	r := gin.Default()

	user.RegisterRoutes(r)

	r.Run(":8080") // http://localhost:8080
}