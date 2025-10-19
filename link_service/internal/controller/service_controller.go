package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceController struct{}

func NewServiceController() *ServiceController {
	return &ServiceController{}
}

func (sc *ServiceController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Healthy"})
}
