package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

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

	ServerID, ServerName, err := s.repo.CreateServer(ctx, req.ServerName, user_id)
	if err != nil {
		err = fmt.Errorf("sql error: %s", err)
		return nil, err
	}

	return &CreateServerRes{
		ServerID:   ServerID,
		ServerName: ServerName,
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
		Servers: res,
	}, nil
}

func (s *service) CreateChannel(ctx context.Context, req *CreateChannelReq, server_id int) (*CreateChannelRes, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ChannelID, ChannelName, err := s.repo.CreateChannel(ctx, req.ChannelName, server_id)
	if err != nil {
		err = fmt.Errorf("sql error: %s", err)
		return nil, err
	}

	return &CreateChannelRes{
		ChannelID:   ChannelID,
		ChannelName: ChannelName,
	}, nil
}

func (s *service) GetChannel(ctx context.Context, server_id int) (*GetChannelRes, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.repo.GetChannel(ctx, server_id)
	if err != nil {
		err = fmt.Errorf("sql error: %s", err)
		return nil, err
	}

	return &GetChannelRes{
		Channels: res,
	}, nil
}

func (s *service) GetMember(ctx context.Context, server_id int) (*GetMemberRes, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.repo.GetMember(ctx, server_id)
	if err != nil {
		err = fmt.Errorf("sql error: %s", err)
		return nil, err
	}

	return &GetMemberRes{
		Members: res,
	}, nil
}

func (s *service) CreateMsg(ctx context.Context, req *CreateMsgReq, gctx *gin.Context) (*CreateMsgRes, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	jwtClaims, err := util.ParseJWT(gctx)
	if err != nil {
		err = fmt.Errorf("parse JWT error: %s", err)
		return nil, err
	}
	user_id := jwtClaims.UserID

	res, err := s.repo.CreateMsg(ctx, req.ChannelID, user_id, req.Time, req.Message)
	if err != nil {
		err = fmt.Errorf("sql error: %s", err)
		return nil, err
	}

	return &CreateMsgRes{
		UserID:  res.UserID,
		Time:    res.Time,
		Message: res.Message,
	}, nil
}

func (s *service) GetHistorysMsg(ctx context.Context, req *GetHistorysMsgReq) (*GetHistorysMsgRes, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res, err := s.repo.GetHistorysMsg(ctx, req.ChannelID)
	if err != nil {
		err = fmt.Errorf("sql error: %s", err)
		return nil, err
	}

	return &GetHistorysMsgRes{
		Historys: res,
	}, nil
}
