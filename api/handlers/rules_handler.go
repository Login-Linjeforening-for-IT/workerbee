package handlers

import (
	"net/http"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRules(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	sort := c.DefaultQuery("direction", "asc")
	orderBy := c.DefaultQuery("order_by", "id")

	rules, err := h.Services.Rules.GetRules(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rules":       rules,
		"total_count": rules[0].TotalCount,
	})
}

func (h *Handler) GetRule(c *gin.Context) {
	id := c.Param("id")

	rule, err := h.Services.Rules.GetRule(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, rule)
}

func (h *Handler) DeleteRule(c *gin.Context) {
	id := c.Param("id")

	rule, err := h.Services.Rules.DeleteRule(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, rule)
}
