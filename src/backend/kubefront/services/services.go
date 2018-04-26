package services

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
	var services corev1.ServiceList
	s.ctx.Client.List(c, s.ctx.Config.Namespace, &services)
	c.JSON(http.StatusOK, services.Items)
}

func (s *Service) listNamespace(c *gin.Context) {
	var services corev1.ServiceList
	s.ctx.Client.List(c, c.Param("namespace"), &services)
	c.JSON(http.StatusOK, services.Items)
}

func (s *Service) delete(c *gin.Context) {
	id, _ := s.ctx.RunWorker(c.MustGet("username").(string), "service/"+c.Param("namespace")+"/"+c.Param("id"), func(ctx context.Context) (interface{}, error) {
		service := &corev1.Service{
			Metadata: &metav1.ObjectMeta{
				Name:      k8s.String(c.Param("id")),
				Namespace: k8s.String(c.Param("namespace")),
			},
		}
		time.Sleep(time.Second * 10)
		return nil, s.ctx.Client.Delete(s.ctx, service)
	})
	c.JSON(http.StatusOK, gin.H{
		"status":  "JOB_STARTED",
		"jobId":   id,
		"message": "Deletion job started",
	})
}
