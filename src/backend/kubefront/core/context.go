package core

import (
	"context"
	"database/sql"

	"github.com/ericchiang/k8s"
)

//Context is used for kubefront backend states
type Context struct {
	context.Context
	Config    *Config
	Client    *k8s.Client
	Upstreams map[string]map[string]*Upstream
	Workers   map[string]map[string]*Worker
	Database  *sql.DB
}

//Worker is a backgroud job
type Worker struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Entity  string `json:"entity"`
}
