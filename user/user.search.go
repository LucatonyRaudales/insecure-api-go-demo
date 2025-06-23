package user
import (
	"net/http"

    "strings"
	"github.com/gin-gonic/gin"
)
func SearchUserV2(c *gin.Context) {
    q := c.Query("name")

    if len(q) < 2 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "query too short (min 2 characters)"})
        return
    }

    var results []gin.H 
    for _, u := range Users {
        if strings.Contains(strings.ToLower(u.Name), strings.ToLower(q)) {
            results = append(results, gin.H{"id": u.ID, "name": u.Name})
        }
    }

    if len(results) == 0 {
        c.JSON(http.StatusOK, gin.H{"message": "no users found matching your query"})
    } else {
        c.JSON(http.StatusOK, results)
    }
}
