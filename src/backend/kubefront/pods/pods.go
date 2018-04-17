package pods

import (
	"net/http"

	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	"github.com/gin-gonic/gin"
)

func (s *Service) list(c *gin.Context) {
	var pods corev1.PodList
	s.ctx.Client.List(c, s.ctx.Config.Namespace, &pods)
	c.JSON(http.StatusOK, pods.Items)
}
