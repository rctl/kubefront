package deployments

import (
	"context"
	"net/http"
	"time"

	"github.com/ericchiang/k8s"
	"github.com/ericchiang/k8s/apis/extensions/v1beta1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
	"github.com/gin-gonic/gin"
)

func (s *Service) list(c *gin.Context) {
	var deployments v1beta1.DeploymentList
	s.ctx.Client.List(c, s.ctx.Config.Namespace, &deployments)
	c.JSON(http.StatusOK, deployments.Items)
}

func (s *Service) listNamespace(c *gin.Context) {
	var deployments v1beta1.DeploymentList
	s.ctx.Client.List(c, c.Param("namespace"), &deployments)
	c.JSON(http.StatusOK, deployments.Items)
}

func (s *Service) delete(c *gin.Context) {
	id, _ := s.ctx.RunWorker(c.MustGet("username").(string), c.Param("id"), func(ctx context.Context) (interface{}, error) {
		deployment := &v1beta1.Deployment{
			Metadata: &metav1.ObjectMeta{
				Name:      k8s.String(c.Param("id")),
				Namespace: k8s.String(c.Param("namespace")),
			},
		}
		time.Sleep(time.Second * 10)
		return nil, s.ctx.Client.Delete(s.ctx, deployment)
	})
	c.JSON(http.StatusOK, gin.H{
		"status":  "JOB_STARTED",
		"jobId":   id,
		"message": "Deletion job started",
	})
}
