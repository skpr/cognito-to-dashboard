package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *DashboardServer) Readyz(c *gin.Context) {
	c.String(http.StatusOK, "Ready!")
}
