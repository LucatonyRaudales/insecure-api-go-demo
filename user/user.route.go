package user

import (
	"github.com/gin-gonic/gin"
	"insecure-api-go/middlewares"

)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", GetAllUsers)
	r.GET("/search", SearchUser)
	r.POST("/login", LoginUser)
	r.POST("/comment", PostComment)
	r.GET("/comments", GetComments)
}

func RegisterSecureRoutes(r *gin.Engine) {
	r.POST("/v2/login", LoginUserV2)
	v2 := r.Group("/v2", middlewares.AuthMiddleware())
	{
		v2.GET("/users", GetAllUsersV2)
		v2.GET("/search", SearchUserV2)
		v2.POST("/comment", PostCommentV2)
	}
}
