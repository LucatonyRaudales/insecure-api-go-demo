package main

import (
	"github.com/gin-gonic/gin"
	"insecure-api-go/user"
)

func main() {
	r := gin.Default()
	user.Init()

	user.RegisterRoutes(r)
	user.RegisterSecureRoutes(r)

	r.Run(":8080") 
}
