package handlers

import (
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
)

// GetDocs godoc
// @Summary      Get Swagger docs
// @Description  Serves the Swagger UI for API documentation.
// @Tags         docs
// @Produce      html
// @Success      200  {string}  string
// @Failure      500  {object}  error
// @Router       /api/v2/docs [get]
func GetDocs(c *gin.Context) {
	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecURL: "./docs/swagger.json",
		CustomOptions: scalar.CustomOptions{
			PageTitle: "Workerbee",
		},
		DarkMode: true,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Send HTML to client
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
}
