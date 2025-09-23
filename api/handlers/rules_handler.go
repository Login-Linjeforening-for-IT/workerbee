package handlers

import (
	"log"
	"net/http"
	"strings"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

var allowedSortColumnsRules = map[string]string{
	"id":         "r.id",
	"name_no":    "r.name_no",
	"name_en":    "r.name_en",
	"created_at": "r.created_at",
	"updated_at": "r.updated_at",
}

func (h *Handler) GetRules(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	sort := c.DefaultQuery("direction", "asc")
	orderBy := c.DefaultQuery("order_by", "id")

	orderBySanitized, sortSanitized, err := internal.SanitizeSort(orderBy, sort, allowedSortColumnsRules)
	if internal.HandleError(c, err) {
		return
	}

	rules, err := h.Rules.GetRules(search, limit, offset, strings.ToUpper(sortSanitized), orderBySanitized)
	if internal.HandleError(c, err) {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jobs":        rules,
		"total_count": rules[0].TotalCount,
	})
}
