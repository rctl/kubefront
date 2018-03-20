package kubefront

import (
	"github.com/gin-gonic/gin"
	"github.com/rctl/kubefront/src/backend/kubefront/authentication"
	"github.com/rctl/kubefront/src/backend/kubefront/core"
	"k8s.io/client-go/kubernetes"
)

//Server is a kubefront backend server instance
type Server struct {
	core.Context
}

//New creates a new instance of a kubefront server
func New(JWTSectet string, client *kubernetes.Clientset) *Server {
	return &Server{
		Config: &core.Config{
			JWTSecret: JWTSectet,
		},
		Client: client,
	}
}

//Serve starts the Kubefront API and makes it accessable
func (s *Server) Serve(addr []string) error {
	r := gin.Default()
	//Register API routes
	authentication.Routes(r.Group("/auth/"), s)
	//Start server
	return r.Run(addr...)
}
