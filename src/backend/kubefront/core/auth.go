package core

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthMiddleware authenticates a user by its token
func AuthMiddleware(ctx *Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwt.Parse(c.Request.Header.Get("Token"), func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return ctx.Config.JWTSecret, nil
		})

		if err != nil || !token.Valid {
			c.String(http.StatusForbidden, "Supplied token was not valid.")
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("scopes", claims["scopes"])
			c.Set("username", claims["username"])
			c.Next()
			return
		}
		c.String(http.StatusInternalServerError, "Could not decode supplied access token.")
	}
}
