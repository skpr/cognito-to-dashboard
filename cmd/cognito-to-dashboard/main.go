package main

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"

	"github.com/skpr/cognito-to-dashboard/internal/server"
)

func main() {
	config, err := server.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	cfg, err := awsconfig.LoadDefaultConfig(
		ctx,
		awsconfig.WithRegion(config.Region),
		awsconfig.WithCredentialsProvider(aws.AnonymousCredentials{}),
	)
	if err != nil {
		panic(err)
	}

	server := server.NewDashboardServer(
		ctx,
		config,
		cognitoidentity.NewFromConfig(cfg),
		cache.New(5*time.Minute, 10*time.Minute),
	)

	router := gin.Default()

	router.GET("/readyz", server.Readyz)
	router.GET("/goto/:dashboard", server.GoTo)
	router.GET("/callback", server.Callback)

	// https://gin-gonic.com/docs/examples/graceful-restart-or-stop
	// This Getenv is set to mirror what Gin does.
	// https://github.com/gin-gonic/gin/blob/44d0dd70924dd154e3b98bc340accc53484efa9c/utils.go#L143
	err = endless.ListenAndServe(os.Getenv("PORT"), router)
	if err != nil {
		panic(err)
	}
}
