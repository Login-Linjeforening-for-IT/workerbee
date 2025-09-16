package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetForm godoc
// @Summary      Get a single form
// @Description  Returns a form by its ID
// @Tags         forms
// @Param        id   path      string  true  "Form ID"
// @Success      200  {object}  models.Form
// @Failure      500  {object}  error
// @Router       /api/v2/forms/{id} [get]
func (h *Handler) GetForm(c *gin.Context) {
	id := c.Param("id")

	forms, err := h.Forms.GetForm(id)
	if err != nil {

	}

	if len(forms) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
		return
	}

	c.JSON(http.StatusOK, forms[0])
}

// GetForms godoc
// @Summary      List forms
// @Description  Returns a list of forms, supports search and pagination
// @Tags         forms
// @Param        search    query     string  false  "Search string"
// @Param        limit     query     int     false  "Limit"
// @Param        offset    query     int     false  "Offset"
// @Success 200 {object} models.FormsResponse
// @Failure      500  {object}  error
// @Router       /api/v2/forms [get]
func (h *Handler) GetForms(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")

	forms, err := h.Forms.GetForms(search, limit, offset)
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"forms":       forms,
		"total_count": forms[0].TotalCount,
	})
}
