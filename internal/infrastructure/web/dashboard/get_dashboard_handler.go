package dashboard

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"place4live/internal/application/port/in"
)

const path = "/dashboard/:name"
const paramName = "name"

type GetDashboardHandler struct {
	port in.GetCityInPort
}

func NewGetDashboardHandler(port in.GetCityInPort) *GetDashboardHandler {
	return &GetDashboardHandler{port: port}
}

func (gdh *GetDashboardHandler) Handle(ctx *gin.Context) {
	name := ctx.Param(paramName)
	city := <-gdh.port.GetCity(name)

	if city != nil {
		ctx.JSON(http.StatusOK, city)
	} else {
		ctx.Status(http.StatusNotFound)
	}
}

func (gdh *GetDashboardHandler) Path() string {
	return path
}

func (gdh *GetDashboardHandler) Method() string {
	return http.MethodGet
}
