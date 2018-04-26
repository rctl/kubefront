package kubefront

import (
	"context"
	"database/sql"

	"github.com/ericchiang/k8s"
	"github.com/gin-gonic/gin"
	"github.com/rctl/kubefront/src/backend/kubefront/authentication"
	"github.com/rctl/kubefront/src/backend/kubefront/core"
	"github.com/rctl/kubefront/src/backend/kubefront/deployments"
	"github.com/rctl/kubefront/src/backend/kubefront/nodes"
	"github.com/rctl/kubefront/src/backend/kubefront/pods"
	"github.com/rctl/kubefront/src/backend/kubefront/services"
	"github.com/rctl/kubefront/src/backend/kubefront/workers"
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
			Workers:  make(map[string]map[string]*core.Worker),
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
		workers.Routes(r.Group("/workers"), s.Context)
		deployments.Routes(r.Group("/deployments"), s.Context)
		services.Routes(r.Group("/services"), s.Context)
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
