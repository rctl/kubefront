package nodes

import (
	"net/http"

	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	"github.com/gin-gonic/gin"
)

func (s *Service) list(c *gin.Context) {
	var nodes corev1.NodeList
	s.ctx.Client.List(c, s.ctx.Config.Namespace, &nodes)
	c.JSON(http.StatusOK, nodes.Items)
}
