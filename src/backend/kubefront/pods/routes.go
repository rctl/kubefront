package pods

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
		r.GET("/:namespace/", s.listNamespace)
		r.DELETE("/:namespace/:id", s.delete)
	}
	go func(s *Service) {
		//Watch for node changes
		var pod corev1.Pod
		watcher, err := s.ctx.Client.Watch(s.ctx, s.ctx.Config.Namespace, &pod)
		defer watcher.Close()
		if err != nil {
			panic(err.Error())
		}
		for {
			n := new(corev1.Pod)
			watcher.Next(n)
			if n == nil {
				continue
			}
			s.ctx.NotifySubscribers("PODS", &core.Message{
				Action: "POD_CHANGED",
				Entity: *n.Metadata.Name,
				Data:   n,
			})
		}
	}(s)
}
