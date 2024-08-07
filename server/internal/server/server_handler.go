package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) GetServer(ctx *gin.Context) {
	res, err := h.service.GetServer(ctx.Request.Context(), ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
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

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) JoinServer(ctx *gin.Context) {
	var s JoinServerReq
	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.service.JoinServer(ctx.Request.Context(), &s, ctx)
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

func (h *Handler) GetMember(ctx *gin.Context) {
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
