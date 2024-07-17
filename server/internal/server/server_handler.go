package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
	hub     *Hub
}

func NewHandler(s Service, h *Hub) *Handler {
	return &Handler{
		service: s,
		hub:     h,
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

	res, err := h.service.CreateServer(ctx.Request.Context(), &s, ctx)
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

func (h *Handler) GetServerByEmail(ctx *gin.Context) {
	res, err := h.service.GetServerByEmail(ctx.Request.Context(), ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
