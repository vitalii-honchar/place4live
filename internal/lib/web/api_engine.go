package web

import (
	"github.com/gin-gonic/gin"
)

type ApiHandler interface {
	Path() string
	Handle(ctx *gin.Context)
	Method() string
}

type ApiEngine struct {
	engine *gin.Engine
}

func (ae *ApiEngine) AddHandler(h ApiHandler) {
	ae.engine.Handle(h.Method(), h.Path(), func(ctx *gin.Context) {
		h.Handle(ctx)
	})
}

func (ae *ApiEngine) AddMiddleware(f gin.HandlerFunc) {
	ae.engine.Use(f)
}

func (ae *ApiEngine) Run() {
	ae.engine.Run()
}

func NewApiEngine() *ApiEngine {
	return &ApiEngine{gin.New()}
}
