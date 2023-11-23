package oauth2utils

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
)

type LoginParams struct {
	ClientID string
	Callback string
	AuthURL  string
	TokenURL string
	Code     string
}

func (p LoginParams) Validate() error {
	var errs []error

	if p.ClientID == "" {
		errs = append(errs, fmt.Errorf("client id not provided"))
	}

	if p.Callback == "" {
		errs = append(errs, fmt.Errorf("callback not provided"))
	}

	if p.AuthURL == "" {
		errs = append(errs, fmt.Errorf("auth url not provided"))
	}

	if p.TokenURL == "" {
		errs = append(errs, fmt.Errorf("token url not provided"))
	}

	if p.Code == "" {
		errs = append(errs, fmt.Errorf("code not provided"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func Login(ctx context.Context, params LoginParams) (string, error) {
	if err := params.Validate(); err != nil {
		return "", fmt.Errorf("validation failed: %w", err)
	}

	config := oauth2.Config{
		RedirectURL: params.Callback,
		ClientID:    params.ClientID,
		Endpoint: oauth2.Endpoint{
			AuthStyle: oauth2.AuthStyleInParams,
			AuthURL:   params.AuthURL,
			TokenURL:  params.TokenURL,
		},
	}

	token, err := config.Exchange(ctx, params.Code)
	if err != nil {
		return "", fmt.Errorf("failed to login with code: %w", err)
	}

	// Extract the ID Token from OAuth2 token.
	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		return "", fmt.Errorf("missing id_token: %w", err)
	}

	return idToken, nil
}
