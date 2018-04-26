package deployments

import (
	"github.com/ericchiang/k8s/apis/extensions/v1beta1"
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
	//Setup routes with need for authentication
	r.Use(core.AuthMiddleware(ctx))
	{
		r.GET("/", s.list)
		r.GET("/:namespace/", s.listNamespace)
		r.DELETE("/:namespace/:id", s.delete)
	}
	go func(s *Service) {
		//Watch for node changes
		var deployment v1beta1.Deployment
		watcher, err := s.ctx.Client.Watch(s.ctx, s.ctx.Config.Namespace, &deployment)
		defer watcher.Close()
		if err != nil {
			panic(err.Error())
		}
		for {
			n := new(v1beta1.Deployment)
			watcher.Next(n)
			if n == nil {
				continue
			}
			s.ctx.NotifySubscribers("DEPLOYMENTS", &core.Message{
				Action: "DEPLOYMENT_CHANGED",
				Entity: *n.Metadata.Name,
				Data:   n,
			})
		}
	}(s)
}
