package kubefront

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

//InititalizeEmptyDatabase populates an empty database with tables needed for kubefront
func (s *Server) InititalizeEmptyDatabase() error {
	ctx := context.Background()
	defer ctx.Done()
	tx, err := s.Database.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	var count int
	err = tx.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='users';").Scan(&count)
	if err != nil || count == 1 {
		tx.Rollback()
		return errors.New("Database already initialized")
	}
	_, err = tx.Exec(`CREATE TABLE users (
        username VARCHAR(255) NOT NULL,
		password VARCHAR(1024) NOT NULL,
		PRIMARY KEY (username)
	);`)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(`CREATE TABLE sessions (
        username VARCHAR(255) NOT NULL,
		session VARCHAR(255) NOT NULL,
		PRIMARY KEY (username,session)
	);`)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(`CREATE TABLE permissions (
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
	ctx := context.Background()
	defer ctx.Done()
	tx, err := s.Database.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}
	password, err := password.Generate(32, 5, 3, false, false)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	rows, err := tx.Query("SELECT COUNT(username) FROM users WHERE username='admin'")
	defer rows.Close()
	if err != nil {
		tx.Rollback()
		return "", err
	}
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	if count != 0 {
		tx.Rollback()
		return "", errors.New("admin user already exists")
	}
	prep, err := tx.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	defer prep.Close()
	if err != nil {
		tx.Rollback()
		return "", err
	}
	_, err = prep.Exec("admin", hashedPassword)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	_, err = tx.Exec("INSERT INTO permissions (username, scope, permission) VALUES ('admin', '.+', '.+')")
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
