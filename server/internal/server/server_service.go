package server

import (
	"context"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"

	"server/util"
)

type service struct {
	timeout time.Duration
	repo    Respository
}

func NewService(repo Respository) *service {
	return &service{
		repo:    repo,
		timeout: 5 * time.Second,
	}
}

func (s *service) CreateServer(ctx context.Context, req *CreateServerReq, gctx *gin.Context) (*CreateServerRes, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// We parse the JWT token from the cookie to get logging user's ID
	jwtClaims, err := util.ParseJWT(gctx)
	if err != nil {
		err = fmt.Errorf("parse JWT error: %s", err)
		return nil, err
	}
	user_id := jwtClaims.UserID

	server := &CreateServerReq{
		ServerName: req.ServerName,
	}

	res, err := s.repo.CreateServer(ctx, server, user_id)
	if err != nil {
		err = fmt.Errorf("sql error: %s", err)
		return nil, err
	}

	return &CreateServerRes{
		ServerId:   res.ServerId,
		ServerName: res.ServerName,
	}, nil
}

func (s *service) GetServerByEmail(ctx context.Context, gctx *gin.Context) (*GetServerRes, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// We parse the JWT token from the cookie to get logging user's ID
	jwtClaims, err := util.ParseJWT(gctx)
	if err != nil {
		err = fmt.Errorf("parse JWT error: %s", err)
		return nil, err
	}
	// fmt.Println("jwtClaims: ", jwtClaims)
	email := jwtClaims.Email

	res, err := s.repo.GetServerByEmail(ctx, email)
	if err != nil {
		err = fmt.Errorf("sql error: %s", err)
		return nil, err
	}

	return &GetServerRes{
		Servers:   res.Servers,
	}, nil
}