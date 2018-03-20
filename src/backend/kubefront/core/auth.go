package core

import (
	"github.com/gin-gonic/gin"
)

//AuthMiddleware authenticates a user by its token
func AuthMiddleware(ctx *Context) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
