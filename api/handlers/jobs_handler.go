package handlers

import (
	"net/http"
	"strings"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

var allowedSortColumnsCities = map[string]string{
	"id":   "c.id",
	"name": "c.name",
}

func (h *Handler) GetJobs(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "desc")

	jobs, err := h.Services.Jobs.GetJobs(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jobs":        jobs,
		"total_count": jobs[0].TotalCount,
	})
}

func (h *Handler) GetJob(c *gin.Context) {
	id := c.Param("id")

	job, err := h.Services.Jobs.GetJob(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, job)
}

func (h *Handler) DeleteJob(c *gin.Context) {
	id := c.Param("id")

	job, err := h.Services.Jobs.DeleteJob(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, job)
}

func (h *Handler) GetCities(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	sort := c.DefaultQuery("direction", "asc")
	orderBy := c.DefaultQuery("order_by", "id")

	orderBySanitized, sortSanitized, err := internal.SanitizeSort(orderBy, sort, allowedSortColumnsCities)
	if internal.HandleError(c, err) {
		return
	}

	cities, err := h.Services.Jobs.GetCities(search, limit, offset, strings.ToUpper(sortSanitized), orderBySanitized)

	c.JSON(http.StatusOK, gin.H{
		"cities":      cities,
		"total_count": cities[0].TotalCount,
	})
}
