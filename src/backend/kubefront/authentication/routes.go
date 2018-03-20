package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/rctl/kubefront/src/backend/kubefront/core"
)

//Service is an instance of the authentication API handler
type Service struct {
	config *core.Config
}

//Routes setup routes for the Authentication API
func Routes(r *gin.RouterGroup, config *core.Config) {
	//MOD: This struct holds a config, which is passed from the implementing package, it contains needed global params for routes to function
	s := Service{
		config: config,
	}
	//Setup routes without need for authentication
	r.POST("/", s.authenticate)
	//Setup routes with need for authentication
	r.Use(core.AuthMiddleware(config))
	{
		r.GET("/", s.profile)
	}
}
