package server

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/skpr/cognito-to-dashboard/internal/random"
)

func (s *DashboardServer) GoTo(c *gin.Context) {
	dashboardName := c.Param("dashboard")

	allowed, err := IsAllowed(s.config.AllowedListPath, dashboardName)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if !allowed {
		c.String(http.StatusForbidden, "dashboard not allowed")
		return
	}

	// Generate a state token to prevent request forgery.
	// Also used as our storage key.
	state := random.String(8)

	s.storage.Set(state, dashboardName, s.config.StorageRetention)

	u := url.URL{
		Scheme: "https",
		Host:   s.config.Host,
		Path:   "/login",
	}

	params := url.Values{}
	params.Add("client_id", s.config.ClientID)
	params.Add("response_type", "code")
	params.Add("scope", strings.Join(s.config.Scope, " "))
	params.Add("state", state)
	params.Add("redirect_uri", s.config.Callback)

	u.RawQuery = params.Encode()

	c.Redirect(http.StatusTemporaryRedirect, u.String())
}
