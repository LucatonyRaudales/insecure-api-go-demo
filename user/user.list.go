package user
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SafeUser struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UsersResponse struct {
	Users []SafeUser `json:"users"`
}

func GetAllUsersV2(c *gin.Context) {
	var safeUsers []SafeUser
	for _, u := range Users {
		safeUsers = append(safeUsers, SafeUser{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}
	c.JSON(http.StatusOK, UsersResponse{Users: safeUsers})
}
