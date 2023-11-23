package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	cloudwatchdashboard "github.com/skpr/cognito-to-dashboard/internal/aws/cloudwatch/dashboard"
	"github.com/skpr/cognito-to-dashboard/internal/aws/cognito/credentials"
	"github.com/skpr/cognito-to-dashboard/internal/aws/federation"
	"github.com/skpr/cognito-to-dashboard/internal/oauth2utils"
)

func (s *DashboardServer) Callback(c *gin.Context) {
	var (
		state = c.Query("state")
		code  = c.Query("code")
	)

	if state == "" {
		c.String(http.StatusBadRequest, "state is required")
		return
	}

	if code == "" {
		c.String(http.StatusBadRequest, "state is required")
		return
	}

	dashboard, ok := s.storage.Get(state)
	if !ok {
		c.String(http.StatusNotFound, "state not found")
		return
	}

	dashboardName := fmt.Sprintf("%s", dashboard)

	allowed, err := IsAllowed(s.config.AllowedListPath, dashboardName)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if !allowed {
		c.String(http.StatusForbidden, "dashboard not allowed")
		return
	}

	token, err := oauth2utils.Login(s.context, oauth2utils.LoginParams{
		ClientID: s.config.ClientID,
		Callback: s.config.Callback,
		AuthURL:  fmt.Sprintf("https://%s/oauth2/authorize", s.config.Host),
		TokenURL: fmt.Sprintf("https://%s/oauth2/token", s.config.Host),
		Code:     code,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	creds, err := credentials.GetTempCredentials(s.context, s.cognito, credentials.GetTempCredentialsParams{
		Token:            token,
		IdentityPool:     s.config.IdentityPool,
		IdentityProvider: s.config.IdentityProvider,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	dashboardURL := cloudwatchdashboard.GetURI(s.config.Region, dashboardName)

	url, err := federation.GetSignInLink(creds, s.config.Host, dashboardURL, s.config.SessionDuration)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, url)
}
