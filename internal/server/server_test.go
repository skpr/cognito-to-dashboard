package server

import (
	"testing"
	"time"
)

type MockStorageClient struct{}

func (m *MockStorageClient) Set(key string, value interface{}, duration time.Duration) {

}

func (m *MockStorageClient) Get(key string) (interface{}, bool) {
	return nil, false
}

func TestNewDashboardServer(t *testing.T) {
	// @todo, Implement this.
}
