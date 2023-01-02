package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"place4live/internal/module/web/app/port"
)

const path = "/login"

type Request struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Handler struct {
	inPort port.LoginInPort
}

func NewHandler(inPort port.LoginInPort) *Handler {
	return &Handler{inPort: inPort}
}

func (lh *Handler) Handle(ctx *gin.Context) {
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if token, ok := lh.inPort.Login(req.Username, req.Password); ok {
		ctx.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "username or password is incorrect."})
	}
}

func (lh *Handler) Path() string {
	return path
}

func (lh *Handler) Method() string {
	return http.MethodPost
}
