package server

import (
	"context"
	"time"

	"github.com/skpr/cognito-to-dashboard/internal/aws/cognito/credentials"
)

type DashboardServer struct {
	context context.Context
	config  Config
	storage StorageClient
	cognito credentials.CognitoClient
}

type StorageClient interface {
	Set(string, interface{}, time.Duration)
	Get(string) (interface{}, bool)
}

func NewDashboardServer(ctx context.Context, config Config, cognito credentials.CognitoClient, storage StorageClient) *DashboardServer {
	return &DashboardServer{
		context: ctx,
		config:  config,
		storage: storage,
		cognito: cognito,
	}
}
