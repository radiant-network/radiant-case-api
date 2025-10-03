package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/radiant-network/radiant-case-api/internal/types"
	"net/http"
)

// CreateCasesBatch godoc
// @Summary Create a batch of cases
// @Description Accepts a JSON array of Case
// @Tags Cases
// @Accept json
// @Produce json
// @Param payload	body		types.ListCases	true	"List Cases"
// @Success 202 {object} types.OperationResponse
// @Failure 400 {object} types.OperationErrorResponse
// @Failure 401
// @Security bearerauth
// @Router /cases/batch [post]
func CreateCasesBatch(c *gin.Context) {
	var batch types.ListCases
	if err := c.ShouldBindJSON(&batch); err != nil {
		batchError := types.OperationError{
			Code:  "invalid_request",
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, types.OperationErrorResponse{Status: "error", Id: "batch_12345", Errors: types.JsonArray[types.OperationError]{batchError}})
		return
	}
	c.JSON(http.StatusAccepted, types.OperationResponse{Status: "in_progress", OperationId: "batch_12345"})
}

// CreateCase godoc
// @Summary Create a case
// @Description Create a single case
// @Tags Cases
// @Accept json
// @Produce json
// @Param payload	body types.Case	true	"Case"
// @Success 201 {object} types.CaseSuccessResponse
// @Success 202 {object} types.OperationResponse
// @Failure 400 {object} types.CaseErrorResponse
// @Failure 401
// @Security bearerauth
// @Router /cases [post]
func CreateCase(c *gin.Context) {
	var aCase types.Case
	if err := c.ShouldBindJSON(&aCase); err != nil {

		responseError := types.CaseErrorResponse{
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, responseError)
		return
	}
	if aCase.SequencingExperiments != nil && len(aCase.SequencingExperiments) != 0 {
		// Simulate async operation for updates that require more processing
		// e.g., adding sequencing experiments that may involve creating samples, tasks, etc.
		c.JSON(http.StatusAccepted, types.OperationResponse{Status: "in_progress", OperationId: "batch_12345"})
		return
	}
	c.JSON(http.StatusCreated, types.CaseSuccessResponse{Id: "12345"})
}

// UpdateCase godoc
// @Summary Update partially one case
// @Description Update a case partial fields by id
// @Tags Cases
// @Accept json
// @Produce json
// @Security bearerauth
// @Param id path string true "case id"
// @Param payload body types.PartialCase true "partial fields"
// @Success 200 {object} types.CaseSuccessResponse
// @Success 202 {object} types.OperationResponse
// @Failure 400 {object} types.CaseErrorResponse
// @Failure 401
// @Router /cases/{id} [patch]
func UpdateCase(c *gin.Context) {
	id := c.Param("id")
	var patch types.PartialCase
	if err := c.ShouldBindJSON(&patch); err != nil {
		responseError := types.CaseErrorResponse{
			Id:    id,
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, responseError)
		return
	}
	if patch.SequencingExperiments != nil && len(patch.SequencingExperiments) != 0 {
		// Simulate async operation for updates that require more processing
		// e.g., adding sequencing experiments that may involve creating samples, tasks, etc.
		c.JSON(http.StatusAccepted, types.OperationResponse{Status: "in_progress", OperationId: "batch_12345"})
		return
	}
	c.JSON(http.StatusOK, types.CaseSuccessResponse{Id: id})

}
