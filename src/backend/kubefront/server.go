package kubefront

import (
	"context"
	"database/sql"

	"github.com/ericchiang/k8s"
	"github.com/gin-gonic/gin"
	"github.com/rctl/kubefront/src/backend/kubefront/authentication"
	"github.com/rctl/kubefront/src/backend/kubefront/core"
)

//Server is a kubefront backend server instance
type Server struct {
	*core.Context
}

//New creates a new instance of a kubefront server
func New(JWTSectet string, client *k8s.Client, database *sql.DB) *Server {
	return &Server{
		Context: &core.Context{
			Context: context.Background(),
			Config: &core.Config{
				JWTSecret: JWTSectet,
			},
			Client:   client,
			Database: database,
		},
	}
}

//Serve starts the Kubefront API and makes it accessable
func (s *Server) Serve(addr ...string) error {
	r := gin.Default()
	//Register API routes
	authentication.Routes(r.Group("/auth/"), s.Context)
	//Start server
	return r.Run(addr...)
}
