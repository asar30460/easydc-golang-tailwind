package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

type Handler struct {
	Service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	var u CreateUserReq
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.Service.CreateUser(ctx.Request.Context(), &u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, res)
}

func (h *Handler) Login(ctx *gin.Context) {
	var user LoginUserReq
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.Service.Login(ctx.Request.Context(), &user)
	if err != nil {
		switch err.Error() {
		case "this email hasn't been registered":
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "此電子郵件尚未註冊，現在進行註冊"})

		case "invalid password":
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "密碼錯誤"})

		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
		}
		return
	}

	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("jwt", res.accessToken, 3600, "/", os.Getenv("CLIENT_DOMAIN"), true, false)
	ctx.SetCookie("user_id", strconv.Itoa(int(res.UserID)), 3600, "/", os.Getenv("CLIENT_DOMAIN"), true, false)

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "/", os.Getenv("CLIENT_DOMAIN"), true, false)
	ctx.SetCookie("user_id", "", -1, "/", os.Getenv("CLIENT_DOMAIN"), true, false)
	ctx.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
