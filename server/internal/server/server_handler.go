package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo Respository
	hub *Hub
}

func NewHandler(r Respository, h *Hub) *Handler {
	return &Handler{
		repo: r,
		hub: h,
	}
}

func (h *Handler) CreateServer(ctx *gin.Context) {
	var s CreateServerReq
	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.repo.CreateServer(ctx.Request.Context(), &s, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 資料庫新增成功後，加入 hub
	h.hub.Servers[string(rune(res.ServerId))] = &Server{
		ID:       string(rune(res.ServerId)),
		Name:     res.ServerName,
		Clients:  make(map[string]*Client),
		Channels: make(map[string]*Channel),
	}

	ctx.JSON(http.StatusOK, res)
}

// func (h *Handler) GetServerByEmail(ctx *gin.Context) {
// 	servers := make([]GetServerRes, 0)

// 	for _, server := range h {
// 		servers = append(servers, GetServerRes{
// 			ServerId:   server,
// 			ServerName: server.Name,
// 		})
// 	}

// 	ctx.JSON(http.StatusOK, servers)
// }