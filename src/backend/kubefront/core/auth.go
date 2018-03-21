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
				fmt.Println(fmt.Errorf("Unexpected signing method: %v", token.Header["alg"]))
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(ctx.Config.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.String(http.StatusForbidden, "Supplied token was not valid.")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			session := claims["session"]
			username := claims["username"]
			var count int
			err := ctx.Database.QueryRowContext(c, "SELECT COUNT(*) FROM sessions WHERE session=? AND username=?", session, username).Scan(&count)
			if err != nil || count != 1 {
				c.String(http.StatusForbidden, "Your session has expired. Please sign in again.")
				c.Abort()
				return
			}
			c.Set("scopes", claims["scopes"])
			c.Set("username", username)
			c.Set("session", session)
			c.Next()
			return
		}
		c.String(http.StatusInternalServerError, "Could not decode supplied access token.")
		c.Abort()
	}
}
