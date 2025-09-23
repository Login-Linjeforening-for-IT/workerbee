package handlers

import (
	"log"
	"net/http"
	"strings"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

var allowedSortColumnsJobs = map[string]string{
	"id":           "ja.id",
	"visible":      "ja.visible",
	"highlight":    "ja.highlight",
	"title_no":      "ja.title_no",
	"title_en":      "ja.title_en",
	"job_type":     "ja.job_type",
	"time_expire":  "ja.time_expire",
	"time_publish": "ja.time_publish",
	"created_at":   "ja.created_at",
	"updated_at":   "ja.updated_at",
}

func (h *Handler) GetJobs(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	sort := c.DefaultQuery("direction", "asc")
	orderBy := c.DefaultQuery("order_by", "id")

	orderBySanitized, sortSanitized, err := internal.SanitizeSort(orderBy, sort, allowedSortColumnsJobs)
	if internal.HandleError(c, err) {
		return
	}

	jobs, err := h.Jobs.GetJobs(search, limit, offset, strings.ToUpper(sortSanitized), orderBySanitized)
	if internal.HandleError(c, err) {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, jobs)
}
