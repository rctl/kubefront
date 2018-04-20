package pods

import (
	"context"
	"net/http"
	"time"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
	"github.com/gin-gonic/gin"
)

func (s *Service) list(c *gin.Context) {
	var pods corev1.PodList
	s.ctx.Client.List(c, s.ctx.Config.Namespace, &pods)
	c.JSON(http.StatusOK, pods.Items)
}

func (s *Service) listNamespace(c *gin.Context) {
	var pods corev1.PodList
	s.ctx.Client.List(c, c.Param("namespace"), &pods)
	c.JSON(http.StatusOK, pods.Items)
}

func (s *Service) delete(c *gin.Context) {
	id, _ := s.ctx.RunWorker(c.MustGet("username").(string), c.Param("id"), func(ctx context.Context) (interface{}, error) {
		pod := &corev1.Pod{
			Metadata: &metav1.ObjectMeta{
				Name:      k8s.String(c.Param("id")),
				Namespace: k8s.String(c.Param("namespace")),
			},
		}
		time.Sleep(time.Minute)
		return nil, s.ctx.Client.Delete(s.ctx, pod)
	})
	c.JSON(http.StatusOK, gin.H{
		"status":  "JOB_STARTED",
		"jobId":   id,
		"message": "Deletion job started",
	})
}
