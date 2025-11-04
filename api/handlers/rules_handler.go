package handlers

import (
	"net/http"
	"workerbee/internal"
	"workerbee/models"

	"github.com/gin-gonic/gin"
)

// CreateRule godoc
// @Summary      Create a new rule
// @Description  Creates a new rule with the provided details. Requires authentication.
// @Tags         rules
// @Accept       json
// @Produce      json
// @Param        rule  body      models.Rule  true  "Rule to create"
// @Success      201    {object}  models.Rule
// @Failure      400    {object}  error
// @Failure      500    {object}  error
// @Router       /api/v2/rules [post]
func (h *Handler) CreateRule(c *gin.Context) {
	var rule models.Rule

	if err := c.ShouldBindBodyWithJSON(&rule); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, rule, *h.Services.Validate) {
		return
	}

	ruleResponse, err := h.Services.Rules.CreateRule(rule)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusCreated, ruleResponse)
}

// UpdateRule godoc
// @Summary      Update an existing rule
// @Description  Updates an existing rule with the provided details. Requires authentication.
// @Tags         rules
// @Accept       json
// @Produce      json
// @Param        id     path      string          true  "Rule ID"
// @Param        rule  body      models.Rule  true  "Updated rule details"
// @Success      200    {object}  models.Rule
// @Failure      400    {object}  error
// @Failure      500    {object}  error
// @Router       /api/v2/rules/{id} [put]
func (h *Handler) UpdateRule(c *gin.Context) {
	var rule models.Rule
	id := c.Param("id")

	if err := c.ShouldBindBodyWithJSON(&rule); internal.HandleError(c, err) {
		return
	}

	if internal.HandleValidationError(c, rule, *h.Services.Validate) {
		return
	}

	ruleResponse, err := h.Services.Rules.UpdateRule(id, rule)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, ruleResponse)
}

// GetRuleNames godoc
// @Summary      Get rule names
// @Description  Retrieves a list of all rule names.
// @Tags         rules
// @Produce      json
// @Success      200  {array}   string
// @Failure      500  {object}  error
// @Router       /api/v2/rules/names [get]	
func (h *Handler) GetRuleNames(c *gin.Context) {
	ruleNames, err := h.Services.Rules.GetRuleNames()
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, ruleNames)
}

// GetRules godoc
// @Summary      Get rules
// @Description  Retrieves a list of rules with optional search, pagination, and sorting.
// @Tags         rules
// @Produce      json
// @Param        search    query     string  false  "Search term"
// @Param        limit     query     string  false  "Number of items to return"  default(20)
// @Param        offset    query     string  false  "Number of items to skip"     default(0)
// @Param        order_by  query     string  false  "Field to order by"          default(id)
// @Param        sort      query     string  false  "Sort order (asc or desc)"   default(asc)
// @Success      200       {object}  map[string]interface{}
// @Failure      500       {object}  error
// @Router       /api/v2/rules [get]
func (h *Handler) GetRules(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limit := c.DefaultQuery("limit", "20")
	offset := c.DefaultQuery("offset", "0")
	orderBy := c.DefaultQuery("order_by", "id")
	sort := c.DefaultQuery("sort", "asc")

	rules, err := h.Services.Rules.GetRules(search, limit, offset, orderBy, sort)
	if internal.HandleError(c, err) {
		return
	}

	if len(rules) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"rules":       rules,
			"total_count": 0,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"rules":       rules,
			"total_count": rules[0].TotalCount,
		})
	}
}

// GetRule godoc
// @Summary      Get rule by ID
// @Description  Retrieves a specific rule by its ID.
// @Tags         rules
// @Produce      json
// @Param        id   path      string  true  "Rule ID"
// @Success      200  {object}  models.Rule
// @Failure      500  {object}  error
// @Router       /api/v2/rules/{id} [get]
func (h *Handler) GetRule(c *gin.Context) {
	id := c.Param("id")

	rule, err := h.Services.Rules.GetRule(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, rule)
}

func (h *Handler) DeleteRule(c *gin.Context) {
	id := c.Param("id")

	ruleId, err := h.Services.Rules.DeleteRule(id)
	if internal.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": ruleId})
}
