package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"place4live/internal/application/port/in"
	"place4live/internal/domain"
	"strconv"
)

const path = "/dashboard/:id"
const paramId = "id"

type GetDashboardHandler struct {
	port in.GetCityInPort
}

func NewGetDashboardHandler(port in.GetCityInPort) *GetDashboardHandler {
	return &GetDashboardHandler{port: port}
}

func (gdh *GetDashboardHandler) Handle(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param(paramId), 10, 64)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	} else {
		dashboard := getDashboard(id)
		if dashboard != nil {
			ctx.JSON(http.StatusOK, newUiDashboard(dashboard))
		} else {
			ctx.Status(http.StatusNotFound)
		}
	}
}

func (gdh *GetDashboardHandler) Path() string {
	return path
}

func (gdh *GetDashboardHandler) Method() string {
	return http.MethodGet
}

func getDashboard(id int64) *domain.Dashboard {
	return &domain.Dashboard{
		Id: id,
		Cities: map[int64]*domain.DashboardCity{
			1: {Order: 10, City: domain.City{Name: "Toronto"}},
			2: {Order: 7, City: domain.City{Name: "Calgary"}},
			3: {Order: 5, City: domain.City{Name: "Kyiv"}},
			4: {Order: 23, City: domain.City{Name: "Edmonton"}},
		},
	}
}
