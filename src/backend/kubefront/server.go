package kubefront

import (
	"github.com/gin-gonic/gin"
	"github.com/rctl/kubefront/src/backend/kubefront/authentication"
	"github.com/rctl/kubefront/src/backend/kubefront/data"
)

//Server is a kubefront backend server instance
type Server struct {
	config *data.Config
}

//New creates a new instance of a kubefront server
func New(JWTSectet string) *Server {
	return &Server{
		config: &data.Config{
			JWTSecret: JWTSectet,
		},
	}
}

//Serve starts the Kubefront API and makes it accessable
func (s *Server) Serve(addr []string) error {
	r := gin.Default()
	//Register API routes
	authentication.Routes(r.Group("/auth/"), s.config)
	//Start server
	return r.Run(addr...)
}
