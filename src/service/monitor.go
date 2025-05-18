package service

import (
	"gin_example/src/service/repo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MonitorService struct {
	Repo repo.MonitorRepo
}

func (ms *MonitorService) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Health": ms.Repo.HealthCheck(),
	})
}
