package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSubmission godoc
// @Summary      Get a single submission
// @Description  Returns a submission by its ID
// @Tags         submissions
// @Param        id   path      string  true  "Submission ID"
// @Success      200  {object}  models.Submission
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /api/v2/forms/submissions/{id} [get]
func (h *Handler) GetSubmission(c *gin.Context) {
	formID := c.Param("id")
	submissionID := c.Param("submission_id")

	submission, err := h.Services.Submissions.GetSubmission(formID, submissionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if submission == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Submission not found"})
		return
	}

	c.JSON(http.StatusOK, submission)
}
