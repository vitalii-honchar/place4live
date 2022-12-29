package web

import (
	"github.com/gin-gonic/gin"
)

type ApiHandler interface {
	Path() string
	Handle(ctx *gin.Context)
	Method() string
}

func NewApiEngine(handlers []ApiHandler) *gin.Engine {
	engine := gin.New()

	for _, h := range handlers {
		engine.Handle(h.Method(), h.Path(), func(ctx *gin.Context) {
			h.Handle(ctx)
		})
	}

	return engine
}
