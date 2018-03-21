package core

import (
	"fmt"
	"net/http"
	"regexp"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthMiddleware authenticates a user by its token
func AuthMiddleware(ctx *Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Check token validity
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

		//Handle token claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			//Check session validity
			session := claims["session"]
			username := claims["username"]
			var count int
			err := ctx.Database.QueryRowContext(c, "SELECT COUNT(*) FROM sessions WHERE session=? AND username=?", session, username).Scan(&count)
			if err != nil || count != 1 {
				c.String(http.StatusForbidden, "Your session has expired. Please sign in again.")
				c.Abort()
				return
			}
			//Fetch user permissions
			rows, err := ctx.Database.QueryContext(c, "SELECT scope, permission FROM permissions WHERE username=?", username)
			if err != nil {
				fmt.Println(err.Error())
				c.String(http.StatusForbidden, "Could not fetch user permissions for supplied token.")
				return
			}
			var scope string
			var permission string
			scopes := make(map[string]string)
			for rows.Next() {
				rows.Scan(&scope, &permission)
				scopes[scope] = permission
			}

			//Set scope variables
			c.Set("username", username)
			c.Set("session", session)
			c.Set("scopes", scopes)

			//Check if request matches users permissions
			requestedScope := c.Request.URL.Path[1:]
			requestedPermission := "READ"
			if c.Request.Method == "POST" {
				requestedPermission = "ADD"
			} else if c.Request.Method == "PUT" {
				requestedPermission = "EDIT"
			} else if c.Request.Method == "DELETE" {
				requestedPermission = "REMOVE"
			}
			for s, p := range scopes {
				matchScope, _ := regexp.MatchString(s, requestedScope)
				matchPermission, _ := regexp.MatchString(p, requestedPermission)
				if matchScope && matchPermission {
					//Match in permission and scope allows the request to proceed
					c.Next()
					return
				}
			}

			c.String(http.StatusForbidden, "You do not have permission to perform this action.")
			c.Abort()
			return
		}
		c.String(http.StatusInternalServerError, "Could not decode supplied access token.")
		c.Abort()
	}
}

//CORSMiddleware sets headers for CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
