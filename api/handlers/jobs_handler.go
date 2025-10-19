package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateJob(c *gin.Context) {
	var job models.NewJob

	if err := c.ShouldBindBodyWithJSON(&job); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, job, *h.Services.Validate) {
		return
	}

	jobResponse, err := h.Services.Jobs.CreateJob(job)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, jobResponse)
}

func (h *Handler) GetJobs(c *gin.Context) {
	jobTypes := c.DefaultQuery("jobtypes", "")
	skills := c.DefaultQuery("skills", "")
	cities := c.DefaultQuery("cities", "")
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")

	jobs, err := h.Services.Jobs.GetJobs(search, limit, offset, orderBy, sort, jobTypes, skills, cities)
	if internal.HandleError(c, err) {
		return
	}

	if len(jobs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"jobs":        jobs,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"jobs":        jobs,
			"total_count": jobs[0].TotalCount,
		})
	}
}

func (h *Handler) GetProtectedJobs(c *gin.Context) {
	jobTypes := c.DefaultQuery("jobtypes", "")
	skills := c.DefaultQuery("skills", "")
	cities := c.DefaultQuery("cities", "")
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")

	jobs, err := h.Services.Jobs.GetProtectedJobs(search, limit, offset, orderBy, sort, jobTypes, skills, cities)
	if internal.HandleError(c, err) {
		return
	}

	if len(jobs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"jobs":        jobs,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"jobs":        jobs,
			"total_count": jobs[0].TotalCount,
		})
	}
}

func (h *Handler) GetJob(c *gin.Context) {
	id := c.Param("id")

	job, err := h.Services.Jobs.GetJob(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, job)
}

func (h *Handler) GetProtectedJob(c *gin.Context) {
	id := c.Param("id")

	job, err := h.Services.Jobs.GetJobProtected(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, job)
}

func (h *Handler) UpdateJob(c *gin.Context) {
	id := c.Param("id")
	var job models.NewJob

	if err := c.ShouldBindBodyWithJSON(&job); internal.HandleError(c, err) {
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

	jobId, err := h.Services.Jobs.DeleteJob(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": jobId})
}

func (h *Handler) GetCities(c *gin.Context) {
	cities, err := h.Services.Jobs.GetJobsCities()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, cities)
}

func (h *Handler) GetActiveJobTypes(c *gin.Context) {
	jobTypes, err := h.Services.Jobs.GetJobTypes()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, jobTypes)
}

func (h *Handler) GetJobSkills(c *gin.Context) {
	jobSkills, err := h.Services.Jobs.GetJobSkills()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, jobSkills)
}

func (h *Handler) GetAllJobTypes(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")

	jobTypes, err := h.Services.Jobs.GetAllJobTypes(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		return
	}

	if len(jobTypes) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"job_types":   jobTypes,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"job_types":   jobTypes,
			"total_count": jobTypes[0].TotalCount,
		})
	}
}

func (h *Handler) GetJobType(c *gin.Context) {
	id := c.Param("id")

	jobType, err := h.Services.Jobs.GetOneJobType(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, jobType)
}

func (h *Handler) CreateJobType(c *gin.Context) {
	var jobType models.JobType

	if err := c.ShouldBindBodyWithJSON(&jobType); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, jobType, *h.Services.Validate) {
		return
	}

	jobTypeResponse, err := h.Services.Jobs.CreateJobType(jobType)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, jobTypeResponse)
}

func (h *Handler) UpdateJobType(c *gin.Context) {
	var jobType models.JobType
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&jobType); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, jobType, *h.Services.Validate) {
		return
	}

	jobTypeResponse, err := h.Services.Jobs.UpdateJobType(id, jobType)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, jobTypeResponse)
}

func (h *Handler) DeleteJobType(c *gin.Context) {
	id := c.Param("id")

	jobTypeId, err := h.Services.Jobs.DeleteJobType(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": jobTypeId})
}
