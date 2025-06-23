package user
import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllUsersV2(c *gin.Context) {
	var safeUsers []gin.H
	for _, u := range Users {
		safeUsers = append(safeUsers, gin.H{
			"id":    u.ID,
			"name":  u.Name,
			"email": u.Email,
		})
	}
	c.JSON(http.StatusOK, safeUsers)
}

func SearchUserV2(c *gin.Context) {
	q := c.Query("name")
	if len(q) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query too short"})
		return
	}
	var results []gin.H
	for _, u := range Users {
		if strings.Contains(strings.ToLower(u.Name), strings.ToLower(q)) {
			results = append(results, gin.H{"id": u.ID, "name": u.Name})
		}
	}
	c.JSON(http.StatusOK, results)
}

func LoginUserV2(c *gin.Context) {
	type Credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	for _, u := range Users {
		if u.Email == creds.Email && u.Password == creds.Password {
			c.JSON(http.StatusOK, gin.H{"token": u.Token})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
}

func PostCommentV2(c *gin.Context) {
	type CommentInput struct {
		Comment string `json:"comment"`
	}
	var input CommentInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment"})
		return
	}
	clean := strings.ReplaceAll(input.Comment, "<", "&lt;")
	clean = strings.ReplaceAll(clean, ">", "&gt;")
	Comments = append(Comments, clean)
	c.JSON(http.StatusOK, gin.H{"message": "comment saved"})
}