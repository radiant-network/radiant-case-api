package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/radiant-network/radiant-case-api/internal/types"
	"net/http"
	"strings"
)

// GetOperation godoc
// @Summary Get status of an asynchronous operation
// @Description Get status of an asynchronous operation
// @Tags Operations
// @Produce json
// @Param id path string true "operation id"
// @Success 201 {object} types.OperationResponse
// @Failure 400 {object} types.OperationErrorResponse
// @Failure 401
// @Security bearerauth
// @Router /operations/{id} [get]
func GetOperation(c *gin.Context) {
	id := c.Param("id")
	if strings.HasPrefix(id, "error-patient-") {
		batchErrors := types.JsonArray[types.OperationError]{
			types.OperationError{Field: "patient[0]", Error: "Patient with organization_patient_id 123 and organization_code CHOP already exist", Code: "PATIENT_ALREADY_EXIST"},
		}
		c.JSON(http.StatusBadRequest, types.OperationErrorResponse{Status: "error", Id: "batch_12345", Errors: batchErrors})
		return
	}
	if strings.HasPrefix(id, "error-sample-") {
		batchErrors := types.JsonArray[types.OperationError]{
			types.OperationError{Field: "sample[0]", Error: "Patient with organization_patient_id 123 and organization_code CHOP does not exist", Code: "PATIENT_NOT_FOUND"},
			types.OperationError{Field: "sample[1]", Error: "Parent sample with sumbitter id 1234 dos not exist", Code: "PARENT_SAMPLE_NOT_FOUND"},
		}
		c.JSON(http.StatusBadRequest, types.OperationErrorResponse{Status: "error", Id: "batch_12345", Errors: batchErrors})
		return
	}
	if strings.HasPrefix(id, "error-case-") {
		batchErrors := types.JsonArray[types.OperationError]{
			types.OperationError{Field: "case[0].patient[0]", Error: "Patient with organization_patient_id 123 and organization_code CHOP not found", Code: "PATIENT_NOT_FOUND"},
			types.OperationError{Field: "case[0].sequencing_experiment[0]", Error: "Sample with submitter_sample_id 123 not found", Code: "SAMPLE_NOT_FOUND"},
			types.OperationError{Field: "case[0].sequencing_experiment[1].task[3]", Error: "Task type with code unknown not supported", Code: "BAD_TASK_TYPE"},
		}
		c.JSON(http.StatusBadRequest, types.OperationErrorResponse{Status: "error", Id: "batch_12345", Errors: batchErrors})
		return
	}
	c.JSON(http.StatusOK, types.OperationResponse{Status: "complete", OperationId: id})
}
