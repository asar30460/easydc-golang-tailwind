package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateServerReq struct {
	ServerName string `json:"server_name"`
}

type CreateServerRes struct {
	ServerID   int    `json:"server_id"`
	ServerName string `json:"server_name"`
}

type GetServerRes struct {
	Servers map[string]string `json:"servers"`
}

type CreateChannelReq struct {
	ChannelName string `json:"channel_name"`
}

type CreateChannelRes struct {
	ChannelID   int    `json:"channel_id"`
	ChannelName string `json:"channel_name"`
}

type GetChannelRes struct {
	Channels map[string]string `json:"channels"`
}

type GetMemberRes struct {
	Members map[string]string `json:"members"`
}

type Msg struct {
	UserID   int
	UserName string
	Time     time.Time
	Message  string
}

type GetHistorysMsgReq struct {
	ChannelID int `json:"channel_id,string,omitempty"`
}

type GetHistorysMsgRes struct {
	Historys []Msg `json:"history_msgs"`
}

type CreateMsgReq struct {
	ChannelID int       `json:"channel_id"`
	Message   string    `json:"msg"`
	Time      time.Time `json:"time"`
}

type CreateMsgRes struct {
	UserID  int       `json:"user_id"`
	Time    time.Time `json:"time"`
	Message string    `json:"msg"`
}

type Service interface {
	CreateServer(ctx context.Context, req *CreateServerReq, gctx *gin.Context) (*CreateServerRes, error)
	GetServerByEmail(ctx context.Context, gctx *gin.Context) (*GetServerRes, error)
	CreateChannel(ctx context.Context, req *CreateChannelReq, server_id int) (*CreateChannelRes, error)
	GetChannel(ctx context.Context, server_id int) (*GetChannelRes, error)
	GetMember(ctx context.Context, server_id int) (*GetMemberRes, error)
	CreateMsg(ctx context.Context, req *CreateMsgReq, gctx *gin.Context) (*CreateMsgRes, error)
	GetHistorysMsg(ctx context.Context, req *GetHistorysMsgReq) (*GetHistorysMsgRes, error)
}

type Respository interface {
	CreateServer(ctx context.Context, server_name string, creator int) (int, string, error)
	GetServerByEmail(ctx context.Context, email string) (map[string]string, error)
	CreateChannel(ctx context.Context, channel_name string, server_id int) (int, string, error)
	GetChannel(ctx context.Context, server_id int) (map[string]string, error)
	GetMember(ctx context.Context, server_id int) (map[string]string, error)
	CreateMsg(ctx context.Context, channel_id int, user_id int, time time.Time, message string) (Msg, error)
	GetHistorysMsg(ctx context.Context, channel_id int) ([]Msg, error)
}
