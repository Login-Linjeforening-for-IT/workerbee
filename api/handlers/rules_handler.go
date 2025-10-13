package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRule(c *gin.Context) {
	var rule models.Rule

	if err := c.ShouldBindBodyWithJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}

	if internal.HandleValidationError(c, rule, *h.Services.Validate) {
		return
	}

	ruleResponse, err := h.Services.Rules.CreateRule(rule)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, ruleResponse)
}

func (h *Handler) UpdateRule(c *gin.Context) {
	var rule models.Rule
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&rule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}

	if internal.HandleValidationError(c, rule, *h.Services.Validate) {
		return
	}

	ruleResponse, err := h.Services.Rules.UpdateRule(id, rule)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, ruleResponse)
}

func (h *Handler) GetRules(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")

	rules, err := h.Services.Rules.GetRules(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		return
	}

	if len(rules) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"rules":       rules,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"rules":       rules,
			"total_count": rules[0].TotalCount,
		})
	}
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
