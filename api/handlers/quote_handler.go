package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateQuote(c *gin.Context) {
	var quotes models.BaseQuote

	if err := c.ShouldBindBodyWithJSON(&quotes); internal.HandleError(c, err) {
		return
	}

	if err := h.Services.Validate.Struct(&quotes); internal.HandleError(c, err) {
		return
	}

	createdQuote, err := h.Services.Quotes.CreateQuote(quotes)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, createdQuote)
}

func (h *Handler) GetQuotes(c *gin.Context) {
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")

	quotes, err := h.Services.Quotes.GetQuotes(limit, offset)
	if internal.HandleError(c, err) {
		return
	}

	if len(quotes) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"quotes":      quotes,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"quotes":      quotes,
			"total_count": quotes[0].TotalCount,
		})
	}
}

func (h *Handler) DeleteQuote(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("user_id")
	admin := c.GetBool("admin")

	deletedQuoteID, err := h.Services.Quotes.DeleteQuote(id, userID, admin)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": deletedQuoteID})
}
