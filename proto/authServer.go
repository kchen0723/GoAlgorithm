package main

import (
	"context"
	"errors"
	// authpb "go-rpc-demo/proto"
)

type AuthService struct {
	UnimplementedAuthServiceServer
}

func (a *AuthService) Auth(c context.Context, req AuthRequest) (*AuthResponse, error) {
	if req.Name != "admin" && req.Password != "123456" {
		return &AuthResponse{Reply: "wrong auth params"}, errors.New("wrong auth params")
	}
	return &AuthResponse{Reply: "Auth pass"}, nil
}
