package handlers

import (
	"net/http"
	"workerbee/models"

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

	form, err := h.Services.Forms.GetForm(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	if form == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
		return
	}
	
	c.JSON(http.StatusOK, form)
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
	orderBy := c.DefaultQuery("order_by", "created_at")
	sort := c.DefaultQuery("sort", "desc")

	forms, err := h.Services.Forms.GetForms(search, limit, offset, orderBy, sort)
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"forms":       forms,
		"total_count": forms[0].TotalCount,
	})
}

// PostForm godoc
// @Summary      Create a new form
// @Description  Creates a new form
// @Tags         forms
// @Accept       json
// @Produce      json
// @Param        form  body  models.Form  true  "Form object"
// @Success      201   {object}  models.Form
// @Failure      400   {object}  error
// @Router       /api/v2/forms [post]
func (h *Handler) PostForm(c *gin.Context) {
	var form models.Form
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newForm, err := h.Services.Forms.PostForm(form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newForm)
}

// PutForm godoc
// @Summary      Update a form
// @Description  Updates a form by ID
// @Tags         forms
// @Accept       json
// @Produce      json
// @Param        id    path  string       true  "Form ID"
// @Param        form  body  models.Form  true  "Form object"
// @Success      200   {object}  models.Form
// @Failure      400   {object}  error
// @Router       /api/v2/forms/{id} [put]
func (h *Handler) PutForm(c *gin.Context) {
	id := c.Param("id")
	var form models.Form
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedForm, err := h.Services.Forms.PutForm(id, form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedForm)
}

// DeleteForm godoc
// @Summary      Delete a form
// @Description  Deletes a form by ID
// @Tags         forms
// @Param        id   path  string  true  "Form ID"
// @Success      200  {object}  models.Form
// @Failure      404  {object}  error
// @Router       /api/v2/forms/{id} [delete]
func (h *Handler) DeleteForm(c *gin.Context) {
	id := c.Param("id")
	deletedForm, err := h.Services.Forms.DeleteForm(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deletedForm)
}