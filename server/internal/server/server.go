package server

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateServerReq struct {
	ServerName string `json:"server_name"`
}

type CreateServerRes struct {
	ServerID   int    `json:"server_id"`
	ServerName string `json:"server_name"`
}

type GetServerRes struct {
	Servers map[int]string `json:"servers"`
}

type JoinServerReq struct {
	ServerID int `json:"server_id"`
}

type JoinServerRes struct {
	ServerID int `json:"server_id"`
	UserID   int `json:"user_id"`
}

type CreateChannelReq struct {
	ChannelName string `json:"channel_name"`
}

type CreateChannelRes struct {
	ChannelID   int    `json:"channel_id"`
	ChannelName string `json:"channel_name"`
}

type GetChannelRes struct {
	Channels map[int]string `json:"channels"`
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
	ChannelID int    `json:"channel_id"`
	Message   string `json:"msg"`
}

type CreateMsgRes struct {
	UserID  int       `json:"user_id"`
	Time    time.Time `json:"time"`
	Message string    `json:"msg"`
}

type Service interface {
	CreateServer(ctx context.Context, req *CreateServerReq, gctx *gin.Context) (*CreateServerRes, error)
	GetServerByEmail(ctx context.Context, gctx *gin.Context) (*GetServerRes, error)
	JoinServer(ctx context.Context, req *JoinServerReq, gctx *gin.Context) (*JoinServerRes, error)
	CreateChannel(ctx context.Context, req *CreateChannelReq, server_id int) (*CreateChannelRes, error)
	GetChannel(ctx context.Context, server_id int) (*GetChannelRes, error)
	GetMember(ctx context.Context, server_id int) (*GetMemberRes, error)
	CreateMsg(ctx context.Context, req *CreateMsgReq, gctx *gin.Context) (*CreateMsgRes, error)
	GetHistorysMsg(ctx context.Context, req *GetHistorysMsgReq) (*GetHistorysMsgRes, error)
}

type Respository interface {
	CreateServer(ctx context.Context, server_name string, creator int) (int, string, error)
	GetServerByEmail(ctx context.Context, email string) (map[int]string, error)
	JoinServer(ctx context.Context, server_id int, user_id int) (int, int, error)
	CreateChannel(ctx context.Context, channel_name string, server_id int) (int, string, error)
	GetChannel(ctx context.Context, server_id int) (map[int]string, error)
	GetMember(ctx context.Context, server_id int) (map[string]string, error)
	CreateMsg(ctx context.Context, channel_id int, user_id int, time time.Time, message string) (Msg, error)
	GetHistorysMsg(ctx context.Context, channel_id int) ([]Msg, error)
}
