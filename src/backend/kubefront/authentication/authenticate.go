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
	tx, err := s.ctx.Database.BeginTx(c, nil)
	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		c.String(http.StatusInternalServerError, "Could not connect to database.")
		return
	}
	//Fetch user password from the database
	var password string
	err = tx.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&password)
	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		c.String(http.StatusForbidden, "Username or password is invalid.")
		return
	}
	//Check if password matches supplied password
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(c.PostForm("password"))); err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		c.String(http.StatusForbidden, "Username or password is invalid.")
		return
	}
	//Create a user session
	session := uuid.Must(uuid.NewV4()).String()
	tx.Exec("INSERT INTO sessions (username,session) VALUES (?, ?)", username, session)
	//Fetch users permissions from database
	rows, err := tx.Query("SELECT scope, permission FROM permissions WHERE username=?", username)
	defer rows.Close()
	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		c.String(http.StatusForbidden, "Could not fetch user permissions for supplied token.")
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		c.String(http.StatusInternalServerError, "Could not setup user session.")
		return
	}
	var scope string
	var permission string
	scopes := make(map[string]string)
	for rows.Next() {
		rows.Scan(&scope, &permission)
		scopes[scope] = permission
	}
	//Create JWT token and pass it to the client
	jwtClaim := jwt.MapClaims{
		"username": username,
		"session":  session,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim).SignedString([]byte(s.ctx.Config.JWTSecret))
	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
		c.String(http.StatusInternalServerError, "Could not generate token on server-side.")
		return
	}
	//Respond with token, session id and the available scopes
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"session": session,
		"scopes":  scopes,
	})
}
