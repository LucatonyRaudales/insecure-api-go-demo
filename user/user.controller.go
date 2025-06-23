package user

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var Comments []string

func GetAllUsers(c *gin.Context) {
	// Exposición de datos sensibles
	c.JSON(http.StatusOK, Users)
}

func SearchUser(c *gin.Context) {
	q := c.Query("name")
	var results []User
	for _, u := range Users {
		if strings.Contains(strings.ToLower(u.Name), strings.ToLower(q)) {
			results = append(results, u)
		}
	}
	c.JSON(http.StatusOK, results)
}

func LoginUser(c *gin.Context) {
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

func PostComment(c *gin.Context) {
	type CommentInput struct {
		Comment string `json:"comment"`
	}
	var input CommentInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment"})
		return
	}

	// No sanitiza => vulnerable a XSS
	Comments = append(Comments, input.Comment)
	c.JSON(http.StatusOK, gin.H{"message": "comment saved"})
}

func GetComments(c *gin.Context) {
	c.JSON(http.StatusOK, Comments)
}