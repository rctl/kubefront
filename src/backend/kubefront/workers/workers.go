package workers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rctl/kubefront/src/backend/kubefront/core"
)

func (s *Service) list(c *gin.Context) {
	l := make([]*core.Worker, 0)
	for _, v := range s.ctx.Workers[c.MustGet("username").(string)] {
		l = append(l, v)
	}
	c.JSON(http.StatusOK, l)
}
