package handlers

import (
	"log"
	"net/http"
	"workerbee/internal"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetJobs(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")

	jobs, err := h.Jobs.GetJobs(search, limit, offset)
	if internal.HandleError(c, err) {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, jobs)
}
