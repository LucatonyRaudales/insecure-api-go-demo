package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", GetAllUsers)
	r.GET("/search", SearchUser)
	r.POST("/login", LoginUser)
	r.POST("/comment", PostComment)
	r.GET("/comments", GetComments)
}
