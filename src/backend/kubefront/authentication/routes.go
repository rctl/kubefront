package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/rctl/kubefront/src/backend/kubefront/core"
)

//Service is an instance of the authentication API handler
type Service struct {
	ctx *core.Context
}

//Routes setup routes for the Authentication API
func Routes(r *gin.RouterGroup, ctx *core.Context) {
	s := &Service{
		ctx: ctx,
	}
	//Setup routes without need for authentication
	r.POST("/", s.authenticate)
	//Setup routes with need for authentication
	r.Use(core.AuthMiddleware(ctx))
	{
		r.GET("/", s.profile)
	}
}
