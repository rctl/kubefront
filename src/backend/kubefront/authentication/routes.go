package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/rctl/kubefront/src/backend/kubefront/data"
)

//Service is an instance of the authentication API handler
type Service struct {
	config *data.Config
}

//Routes setup routes for the Authentication API
func Routes(r *gin.RouterGroup, config *data.Config) {
	//MOD: This struct holds a config, which is passed from the implementing package, it contains needed global params for routes to function
	s := Service{
		config: config,
	}
	//Setup routes without need for authentication
	r.POST("/", s.authenticate)
	//Setup routes with need for authentication
	r.Use(data.AuthMiddleware(config))
	{
		r.GET("/", s.profile)
	}
}
