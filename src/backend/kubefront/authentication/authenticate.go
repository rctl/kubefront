package authentication

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//authenticate exchanges credentials for a valid token
func (s *Service) authenticate(c *gin.Context) {
	username := c.PostForm("username")
	//Fetch user password from the database
	var password string
	err := s.ctx.Database.QueryRowContext(c, "SELECT password FROM users WHERE username=?", username).Scan(&password)
	if err != nil {
		fmt.Println(err.Error())
		c.String(http.StatusForbidden, "Username or password is invalid.")
		return
	}
	//Check if password matches supplied password
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(c.PostForm("password"))); err != nil {
		fmt.Println(err.Error())
		c.String(http.StatusForbidden, "Username or password is invalid.")
		return
	}
	//Get user permissions from database and construct JWT claim
	rows, err := s.ctx.Database.QueryContext(c, "SELECT scope, permission FROM permissions WHERE username=?", username)
	if err != nil {
		fmt.Println(err.Error())
		c.String(http.StatusForbidden, "Username or password is invalid.")
		return
	}
	var scope string
	var permission string
	scopes := make(map[string]string)
	for rows.Next() {
		rows.Scan(&scope, &permission)
		scopes[scope] = permission
	}
	session := uuid.Must(uuid.NewV4()).String()
	tx, err := s.ctx.Database.BeginTx(c, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not connect to backend database.")
		return
	}
	tx.ExecContext(c, "INSERT INTO sessions (username,session) VALUES (?, ?)", username, session)
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		c.String(http.StatusInternalServerError, "Could not setup user session.")
		return
	}
	//Create JWT token and pass it to the client
	jwtClaim := jwt.MapClaims{
		"scopes":   scopes,
		"username": username,
		"session":  session,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim).SignedString([]byte(s.ctx.Config.JWTSecret))
	if err != nil {
		fmt.Println(err.Error())
		c.String(http.StatusInternalServerError, "Could not generate token on server-side.")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"session": session,
		"scopes":  scopes,
	})
}
