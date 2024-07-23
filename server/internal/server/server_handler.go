package server

import (
	"strconv"
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
	h.hub.Servers[res.ServerID] = &Server{
		ID:       res.ServerID,
		Name:     res.ServerName,
		Clients:  make(map[int]*Client),
		Channels: make(map[int]*Channel),
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

func (h *Handler) CreateChannel(ctx *gin.Context) {
	server_id := ctx.Param("server_id")
	int_server_id, _ := strconv.Atoi(server_id)

	var s CreateChannelReq
	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.service.CreateChannel(ctx.Request.Context(), &s, int_server_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetChannel(ctx *gin.Context) {
	server_id := ctx.Param("server_id")
	int_server_id, _ := strconv.Atoi(server_id)

	res, err := h.service.GetChannel(ctx.Request.Context(), int_server_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetMember (ctx *gin.Context) {
	server_id := ctx.Param("server_id")
	int_server_id, _ := strconv.Atoi(server_id)

	res, err := h.service.GetMember(ctx.Request.Context(), int_server_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) CreateMsg(ctx *gin.Context) {
	var s CreateMsgReq
	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.service.CreateMsg(ctx.Request.Context(), &s, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetHistorysMsg(ctx *gin.Context) {
	var s GetHistorysMsgReq
	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.service.GetHistorysMsg(ctx.Request.Context(), &s)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}