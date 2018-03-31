package nodes

import (
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
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
	}
	go func(s *Service) {
		//Watch for node changes
		var node corev1.Node
		watcher, err := s.ctx.Client.Watch(s.ctx, s.ctx.Config.Namespace, &node)
		defer watcher.Close()
		if err != nil {
			panic(err.Error())
		}
		for {
			n := new(corev1.Node)
			watcher.Next(n)
			s.ctx.NotifySubscribers("NODES", &core.Message{
				Action: "NODE_CHANGED",
				Entity: *n.Metadata.Name,
				Data:   n,
			})
		}
	}(s)
}
