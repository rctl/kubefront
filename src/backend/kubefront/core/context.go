package core

import (
	"context"
	"database/sql"

	"github.com/ericchiang/k8s"
)

//Context is used for kubefront backend states
type Context struct {
	context.Context
	Config   *Config
	Client   *k8s.Client
	Database *sql.DB
}
