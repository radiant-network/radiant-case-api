package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/radiant-network/radiant-case-api/internal/types"
	"net/http"
)

// CreatePatientsBatch godoc
// @Summary Create a batch of patients
// @Description Accepts a JSON array of Patients
// @Tags Patients
// @Accept json
// @Produce json
// @Param payload	body		types.ListPatients	true	"List Patients"
// @Success 202 {object} types.OperationResponse
// @Failure 400 {object} types.OperationErrorResponse
// @Failure 401
// @Security bearerauth
// @Router /patients/batch [post]
func CreatePatientsBatch(c *gin.Context) {
	var batch types.ListPatients
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
