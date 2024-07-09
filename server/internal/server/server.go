package server

import (
	"context"
	"github.com/gin-gonic/gin"
)

type CreateServerReq struct {
	ServerName string `json:"server_name"`
}

type CreateServerRes struct {
	ServerId  int `json:"server_id"`
	ServerName string `json:"server_name"`
}

type GetServerRes struct {
	ServerId  int `json:"server_id"`
	ServerName string `json:"server_name"`
}

type Respository interface {
	CreateServer(ctx context.Context, server *CreateServerReq, c *gin.Context) (*CreateServerRes, error)
}