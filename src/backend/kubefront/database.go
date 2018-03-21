package kubefront

import (
	"database/sql"
	"errors"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

//InititalizeEmptyDatabase populates an empty database with tables needed for kubefront
func (s *Server) InititalizeEmptyDatabase() error {
	tx, err := s.Database.BeginTx(s, &sql.TxOptions{})
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(s, `CREATE TABLE users (
        username VARCHAR(255) NOT NULL,
		password VARCHAR(1024) NOT NULL,
		PRIMARY KEY (username)
	);`)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(s, `CREATE TABLE sessions (
        username VARCHAR(255) NOT NULL,
		session VARCHAR(255) NOT NULL,
		PRIMARY KEY (username,session)
	);`)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.ExecContext(s, `CREATE TABLE permissions (
        username VARCHAR(255) NOT NULL,
		scope VARCHAR(64) NOT NULL,
		permission VARCHAR(64) NOT NULL,
		PRIMARY KEY (username,scope),
		CONSTRAINT FK_UserPermission FOREIGN KEY (username) REFERENCES username
	);`)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

//CreateAdminUser creates an admin user and returnes a random password for it
func (s *Server) CreateAdminUser() (string, error) {
	tx, err := s.Database.BeginTx(s, nil)
	if err != nil {
		return "", err
	}
	password, err := password.Generate(32, 5, 3, false, false)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	rows, err := tx.Query("SELECT COUNT(username) FROM users WHERE username='admin'")
	if err != nil {
		return "", err
	}
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	if count != 0 {
		return "", errors.New("admin user already exists")
	}
	prep, err := tx.PrepareContext(s, "INSERT INTO users (username, password) VALUES (?, ?)")
	if err != nil {
		return "", err
	}
	_, err = prep.ExecContext(s, "admin", hashedPassword)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	_, err = tx.ExecContext(s, "INSERT INTO permissions (username, scope, permission) VALUES ('admin', '*', '*')")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return password, nil
}
