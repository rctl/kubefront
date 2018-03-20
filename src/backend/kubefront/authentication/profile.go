package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) profile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"username": c.MustGet("username"),
	})
}
