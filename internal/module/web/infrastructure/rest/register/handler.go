package register

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"place4live/internal/module/web/app/port"
)

const path = "/register"

type Request struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Handler struct {
	inPort port.RegisterInPort
}

func NewHandler(inPort port.RegisterInPort) *Handler {
	return &Handler{inPort: inPort}
}

func (h *Handler) Handle(ctx *gin.Context) {
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.inPort.Register(req.Username, req.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	}
}

func (h *Handler) Path() string {
	return path
}

func (h *Handler) Method() string {
	return http.MethodPost
}
