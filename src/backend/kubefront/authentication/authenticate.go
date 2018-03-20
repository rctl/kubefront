package authentication

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//authenticate exchanges credentials for a valid token
func (s *Service) authenticate(c *gin.Context) {
	username := c.PostForm("username")
	//Fetch user password from the database
	var password string
	err := s.ctx.Database.QueryRowContext(c, "SELECT password FROM users WHERE username='?'", username).Scan(&password)
	if err != nil {
		c.String(http.StatusForbidden, "Username or password is invalid.")
		return
	}
	//Check if password matches supplied password
	if bcrypt.CompareHashAndPassword([]byte(password), []byte(c.PostForm("password"))) != nil {
		c.String(http.StatusForbidden, "Username or password is invalid.")
		return
	}
	//Get user permissions from database and construct JWT claim
	rows, err := s.ctx.Database.QueryContext(c, "SELECT (scope, permission) FROM permissions WHERE username='?'", username)
	var scope string
	var permission string
	scopes := make(map[string]string)
	for rows.Next() {
		rows.Scan(&scope, &permission)
		scopes[scope] = permission
	}
	//Create JWT token and pass it to the client
	jwtClaim := jwt.MapClaims{
		"scopes":   scopes,
		"username": username,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim).SignedString(s.ctx.Config.JWTSecret)
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not generate token on server-side.")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
