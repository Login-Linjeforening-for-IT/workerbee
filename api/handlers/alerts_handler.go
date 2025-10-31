package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAlert(c *gin.Context) {
	var alert models.Alert

	if err := c.ShouldBindBodyWithJSON(&alert); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, alert, *h.Services.Validate) {
		return
	}

	alertResponse, err := h.Services.Alerts.CreateAlert(alert)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, alertResponse)
}

func (h *Handler) GetAllAlerts(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")

	alerts, err := h.Services.Alerts.GetAllAlerts(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, alerts)
}

func (h *Handler) GetAlertByServiceAndPage(c *gin.Context) {
	service := c.Param("service")
	page := c.Param("page")

	alert, err := h.Services.Alerts.GetAlertByServiceAndPage(service, page)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, alert)
}

func (h *Handler) GetAlertByID(c *gin.Context) {
	id := c.Param("id")

	alert, err := h.Services.Alerts.GetAlertByID(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, alert)
}

func (h *Handler) UpdateAlert(c *gin.Context) {
	var alert models.Alert
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&alert); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, alert, *h.Services.Validate) {
		return
	}

	alertResponse, err := h.Services.Alerts.UpdateAlert(id, alert)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, alertResponse)
}

func (h *Handler) DeleteAlert(c *gin.Context) {
	id := c.Param("id")

	deletedID, err := h.Services.Alerts.DeleteAlert(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": deletedID})
}
