package main

import (
	"github.com/gin-gonic/gin"
	"insecure-api-go/user"
)

func main() {
	r := gin.Default()

	// Rutas inseguras
	user.RegisterRoutes(r)
	// Rutas seguras (v2)
	user.RegisterSecureRoutes(r)

	r.Run(":8080") // http://localhost:8080
}
