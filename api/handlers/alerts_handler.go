package handlers

import (
	"net/http"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAlertServices(c *gin.Context) {
	services, err := h.Services.Alerts.GetAlertServices()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, services)
}

func (h *Handler) GetAllPathsInAlertService(c *gin.Context) {
	service := c.Param("service")

	alerts, err := h.Services.Alerts.GetAllPathsInAlertService(service)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, alerts)
}