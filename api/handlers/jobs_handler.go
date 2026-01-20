package handlers

import (
	"fmt"
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

// CreateJob godoc
// @Summary      Create a new job
// @Description  Creates a new job with the provided details.
// @Tags         jobs
// @Accept       json
// @Produce      json
// @Param        job  body      models.NewJob  true  "Job to create"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v2/jobs [post]
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

// GetJobs godoc
// @Summary      Get list of jobs
// @Description  Retrieves a list of jobs with optional filtering, pagination, and sorting.
// @Tags         jobs
// @Produce      json
// @Param        jobtypes  query     string  false  "Filter by job types (comma-separated)"
// @Param        skills    query     string  false  "Filter by skills (comma-separated)"
// @Param        cities    query     string  false  "Filter by cities (comma-separated)"
// @Param        search    query     string  false  "Search term"
// @Param        limit     query     string  false  "Number of results to return"  default(20)
// @Param        offset    query     string  false  "Number of results to skip"    default(0)
// @Param        order_by  query     string  false  "Field to order by"            default(id)
// @Param        sort      query     string  false  "Sort order (asc or desc)"         default(asc)
// @Success      200       {object}  map[string]interface{}
// @Failure      400       {object}  error
// @Failure      500       {object}  error
// @Router       /api/v2/jobs [get]
func (h *Handler) GetJobs(c *gin.Context) {
	jobTypes := c.DefaultQuery("jobtypes", "")
	skills := c.DefaultQuery("skills", "")
	cities := c.DefaultQuery("cities", "")
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")

	jobs, cacheTTL, err := h.Services.Jobs.GetJobs(search, limit, offset, orderBy, sort, jobTypes, skills, cities)
	if internal.HandleError(c, err) {
		return
	}
	c.Header("Cache-Control", fmt.Sprintf("public, max-age=%d", cacheTTL))

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

// GetProtectedJobs godoc
// @Summary      Get list of protected jobs
// @Description  Retrieves a list of protected jobs with optional filtering, pagination, and sorting.
// @Tags         jobs
// @Produce      json
// @Param        jobtypes  query     string  false  "Filter by job types (comma-separated)"
// @Param        skills    query     string  false  "Filter by skills (comma-separated)"
// @Param        cities    query     string  false  "Filter by cities (comma-separated)"
// @Param        search    query     string  false  "Search term"
// @Param        limit     query     string  false  "Number of results to return"  default(20)
// @Param        offset    query     string  false  "Number of results to skip"    default(0)
// @Param        order_by  query     string  false  "Field to order by"            default(id)
// @Param        sort      query     string  false  "Sort order (asc or desc)"         default(asc)
// @Success      200       {object}  map[string]interface{}
// @Failure      400       {object}  error
// @Failure      500       {object}  error
func (h *Handler) GetProtectedJobs(c *gin.Context) {
	jobTypes := c.DefaultQuery("jobtypes", "")
	skills := c.DefaultQuery("skills", "")
	cities := c.DefaultQuery("cities", "")
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")
	historical := c.DefaultQuery("historical", "false")

	jobs, err := h.Services.Jobs.GetProtectedJobs(search, limit, offset, orderBy, sort, jobTypes, skills, cities, historical)
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

// GetJob godoc
// @Summary      Get job by ID
// @Description  Retrieves a specific job by its ID.
// @Tags         jobs
// @Produce      json
// @Param        id   path      string  true  "Job ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v2/jobs/{id} [get]
func (h *Handler) GetJob(c *gin.Context) {
	id := c.Param("id")

	job, err := h.Services.Jobs.GetJob(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, job)
}

// GetProtectedJob godoc
// @Summary      Get protected job by ID
// @Description  Retrieves a specific protected job by its ID.
// @Tags         jobs
// @Produce      json
// @Param        id   path      string  true  "Job ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v2/jobs/protected/{id} [get]
func (h *Handler) GetProtectedJob(c *gin.Context) {
	id := c.Param("id")

	job, err := h.Services.Jobs.GetJobProtected(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, job)
}

// UpdateJob godoc
// @Summary      Update a job
// @Description  Updates a job by ID with the provided details.
// @Tags         jobs
// @Accept       json
// @Produce      json
// @Param        id   path      string        true  "Job ID"
// @Param        job  body      models.NewJob  true  "Job to update"
// @Success      200  {object}  models.NewJob
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v2/jobs/{id} [put]
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

// DeleteJob godoc
// @Summary      Delete a job
// @Description  Deletes a job by its ID.
// @Tags         jobs
// @Produce      json
// @Param        id   path      string  true  "Job ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v2/jobs/{id} [delete]
func (h *Handler) DeleteJob(c *gin.Context) {
	id := c.Param("id")

	jobId, err := h.Services.Jobs.DeleteJob(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": jobId})
}

// GetCities godoc
// @Summary      Get active job cities
// @Description  Retrieves a list of all active job cities.
// @Tags         jobs
// @Produce      json
// @Success      200  {array}   models.Cities
// @Failure      500  {object}  error
// @Router       /api/v2/jobs/cities [get]
func (h *Handler) GetCities(c *gin.Context) {
	cities, err := h.Services.Jobs.GetJobsCities()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, cities)
}

// GetActiveJobTypes godoc
// @Summary      Get active job types
// @Description  Retrieves a list of all active job types.
// @Tags         jobs
// @Produce      json
// @Success      200  {array}   models.JobType
// @Failure      500  {object}  error
// @Router       /api/v2/jobs/types [get]
func (h *Handler) GetActiveJobTypes(c *gin.Context) {
	jobTypes, err := h.Services.Jobs.GetJobTypes()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, jobTypes)
}

// GetJobSkills godoc
// @Summary      Get job skills
// @Description  Retrieves a list of all job skills.
// @Tags         jobs
// @Produce      json
// @Success      200  {array}   models.JobSkills
// @Failure      500  {object}  error
// @Router       /api/v2/jobs/skills [get]
func (h *Handler) GetJobSkills(c *gin.Context) {
	jobSkills, err := h.Services.Jobs.GetJobSkills()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, jobSkills)
}

// GetAllJobTypes godoc
// @Summary      Get all job types
// @Description  Retrieves a list of all job types with optional filtering, pagination, and sorting.
// @Tags         jobs
// @Produce      json
// @Param        search    query     string  false  "Search term"
// @Param        limit     query     string  false  "Number of results to return"  default(20)
// @Param        offset    query     string  false  "Number of results to skip"    default(0)
// @Param        order_by  query     string  false  "Field to order by"            default(id)
// @Param        sort      query     string  false  "Sort order (asc or desc)"         default(asc)
// @Success      200       {object}  map[string]interface{}
// @Failure      400       {object}  error
// @Failure      500       {object}  error
// @Router       /api/v2/jobs/types/all [get]
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

// GetJobType godoc
// @Summary      Get job type by ID
// @Description  Retrieves a specific job type by its ID.
// @Tags         jobs
// @Produce      json
// @Param        id   path      string  true  "Job Type ID"
// @Success      200  {object}  models.JobType
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v2/jobs/types/{id} [get]
func (h *Handler) GetJobType(c *gin.Context) {
	id := c.Param("id")

	jobType, err := h.Services.Jobs.GetOneJobType(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, jobType)
}

// CreateJobType godoc
// @Summary      Create a new job type
// @Description  Creates a new job type with the provided details.
// @Tags         jobs
// @Accept       json
// @Produce      json
// @Param        jobType  body      models.JobType  true  "Job Type to create"
// @Success      201      {object}  models.JobType
// @Failure      400      {object}  error
// @Failure      500      {object}  error
// @Router       /api/v2/jobs/types [post]
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

// UpdateJobType godoc
// @Summary      Update a job type
// @Description  Updates a job type by ID with the provided details.
// @Tags         jobs
// @Accept       json
// @Produce      json
// @Param        id       path      string          true  "Job Type ID"
// @Param        jobType  body      models.JobType  true  "Job Type to update"
// @Success      200      {object}  models.JobType
// @Failure      400      {object}  error
// @Failure      500      {object}  error
// @Router       /api/v2/jobs/types/{id} [put]
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

// DeleteJobType godoc
// @Summary      Delete a job type
// @Description  Deletes a job type by its ID.
// @Tags         jobs
// @Produce      json
// @Param        id   path      string  true  "Job Type ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v2/jobs/types/{id} [delete]
func (h *Handler) DeleteJobType(c *gin.Context) {
	id := c.Param("id")

	jobTypeId, err := h.Services.Jobs.DeleteJobType(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": jobTypeId})
}
