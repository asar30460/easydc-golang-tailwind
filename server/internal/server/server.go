package server

import (
	"context"
	"github.com/gin-gonic/gin"
)

type ServerMetadata struct {
	ServerId   int
	ServerName string
}

type CreateServerReq struct {
	ServerName string `json:"server_name"`
}

type CreateServerRes struct {
	ServerId   int    `json:"server_id"`
	ServerName string `json:"server_name"`
}

type GetServerRes struct {
	Servers     map[string]string `json:"servers"`
}

type Service interface {
	CreateServer(ctx context.Context, req *CreateServerReq, gctx *gin.Context) (*CreateServerRes, error)
	GetServerByEmail(ctx context.Context, gctx *gin.Context) (*GetServerRes, error)
}

type Respository interface {
	CreateServer(ctx context.Context, req *CreateServerReq, creator int) (*ServerMetadata, error)
	GetServerByEmail(ctx context.Context, email string) (*GetServerRes, error)
}
