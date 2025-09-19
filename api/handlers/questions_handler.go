package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.login.no/tekkom/web/beehive/admin-api/v2/models"
)

// PutQuestions godoc
// @Summary      Update questions for a form
// @Description  Accepts an array of questions and updates them
// @Tags         questions
// @Accept       json
// @Produce      json
// @Param        form_id   path      string  true  "Form ID"
// @Param        questions body      []models.Question true "Array of questions"
// @Success      200  {array}  models.Question
// @Failure      400  {object}  error
// @Router       /api/v2/forms/{form_id}/questions [put]
func (h *Handler) PutQuestions(c *gin.Context) {
	formID := c.Param("form_id")
	questions := []models.Question{}

	if err := c.ShouldBindJSON(&questions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedQuestions, err := h.Questions.PutQuestions(formID, questions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, updatedQuestions)
}

// DeleteQuestion godoc
// @Summary      Delete a question
// @Description  Deletes a question by ID
// @Tags         questions
// @Param        id   path  string  true  "Question ID"
// @Success      200  {object}  models.Question
// @Failure      404  {object}  error
// @Router       /api/v2/questions/{id} [delete]
func (h *Handler) DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	
	deletedQuestion, err := h.Questions.DeleteQuestion(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, deletedQuestion)
}
