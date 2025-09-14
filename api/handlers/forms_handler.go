package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/db"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
)

// GetForm godoc
// @Summary      Get a single form
// @Description  Returns a form by its ID
// @Tags         forms
// @Param        id   path      string  true  "Form ID"
// @Success      200  {object}  models.Form
// @Failure      500  {object}  error
// @Router       /api/v2/forms/{id} [get]
func GetForm(c *gin.Context) {
	id := c.Param("id")

	forms := []models.Form{}

	sqlBytes, err := os.ReadFile("./db/forms/get_form.sql")
	if err != nil {
		log.Println("unable to find file, err ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	query := string(sqlBytes)
	err = db.DB.Select(&forms, query, id)
	if err != nil {
		log.Println("unable to query, err ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
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
func GetForms(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")

	forms := []models.FormWithTotalCount{}

	sqlBytes, err := os.ReadFile("./db/forms/get_forms.sql")
	if err != nil {
		log.Println("unable to find file, err ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	query := string(sqlBytes)
	err = db.DB.Select(&forms, query, search, limit, offset)
	if err != nil {
		log.Println("unable to query, err ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	total := 0
	if len(forms) > 0 {
		total = forms[0].TotalCount
	}

	formResponses := make([]models.Form, len(forms))
	for i, f := range forms {
		formResponses[i] = f.Form
	}
	c.JSON(http.StatusOK, gin.H{
		"forms": formResponses,
		"total_count": total,
	})
}