package user
import (
	"log" 
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

var p *bluemonday.Policy

func init() {
    p = bluemonday.StrictPolicy()
    log.Println("bluemonday StrictPolicy initialized in 'user' package.")

}

func PostCommentV2(c *gin.Context) {
    type CommentInput struct {
        Comment string `json:"comment"`
    }
    var input CommentInput
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    sanitizedComment := p.Sanitize(input.Comment)

    Comments = append(Comments, sanitizedComment)

    c.JSON(http.StatusOK, gin.H{"message": "comment saved (sanitized on input)"})
}
