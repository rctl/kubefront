package kubefront

import (
	"context"
	"database/sql"

	"github.com/ericchiang/k8s"
	"github.com/gin-gonic/gin"
	"github.com/rctl/kubefront/src/backend/kubefront/authentication"
	"github.com/rctl/kubefront/src/backend/kubefront/core"
	"github.com/rctl/kubefront/src/backend/kubefront/nodes"
	"github.com/rctl/kubefront/src/backend/kubefront/pods"
)

//Server is a kubefront backend server instance
type Server struct {
	*core.Context
}

//New creates a new instance of a kubefront server
func New(ctx context.Context, JWTSectet string, client *k8s.Client, database *sql.DB) *Server {
	return &Server{
		Context: &core.Context{
			Context: ctx,
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
	r.Use(core.CORSMiddleware())
	{
		//Register API routes
		authentication.Routes(r.Group("/auth/"), s.Context)
		nodes.Routes(r.Group("/nodes"), s.Context)
		pods.Routes(r.Group("/pods"), s.Context)
	}
	r.Use(core.AuthMiddleware(s.Context))
	{
		r.GET("/upstream", func(c *gin.Context) {
			s.upstreamHandler(c)
		})
	}
	//Start server
	return r.Run(addr...)
}
