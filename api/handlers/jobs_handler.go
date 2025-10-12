package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateJob(c *gin.Context) {
	var job models.Job

	if err := c.ShouldBindBodyWithJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}

	if internal.HandleValidationError(c, job, *h.Services.Validate) {
		return
	}

	err := h.Services.Jobs.CreateJob(job)
	if internal.HandleError(c, err) {
		return
	}
}

func (h *Handler) GetJobs(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "desc")

	jobs, err := h.Services.Jobs.GetJobs(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"jobs":        jobs,
		"total_count": jobs[0].TotalCount,
	})
}

func (h *Handler) GetJob(c *gin.Context) {
	id := c.Param("id")

	job, err := h.Services.Jobs.GetJob(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, job)
}

func (h *Handler) UpdateJob(c *gin.Context) {
	id := c.Param("id")
	var job models.Job

	if err := c.ShouldBindBodyWithJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}

	if internal.HandleValidationError(c, job, *h.Services.Validate) {
		return
	}

	jobResponse, err := h.Services.Jobs.UpdateJob(id, job)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, jobResponse)
}

func (h *Handler) DeleteJob(c *gin.Context) {
	id := c.Param("id")

	job, err := h.Services.Jobs.DeleteJob(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, job)
}

func (h *Handler) GetCities(c *gin.Context) {
	cities, err := h.Services.Jobs.GetJobsCities()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cities": cities,
	})
}

func (h *Handler) GetJobTypes(c *gin.Context) {
	jobTypes, err := h.Services.Jobs.GetJobTypes()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"job_types": jobTypes,
	})
}

func (h *Handler) GetJobSkills(c *gin.Context) {
	jobSkills, err := h.Services.Jobs.GetJobSkills()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"job_skills": jobSkills,
	})
}
