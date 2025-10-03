package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/radiant-network/radiant-case-api/internal/types"
)

// CreateSamplesBatch godoc
// @Summary Create a batch of samples
// @Description Accepts a JSON array of Samples
// @Tags Samples
// @Accept json
// @Produce json
// @Param payload	body		types.ListSamples	true	"List Samples"
// @Success 202 {object} types.BatchResponse
// @Failure 400 {object} types.BatchErrorResponse
// @Failure 401
// @Security bearerauth
// @Router /samples/batch [post]
func CreateSamplesBatch(c *gin.Context) {
	var batch types.ListSamples
	if err := c.ShouldBindJSON(&batch); err != nil {
		batchError := types.BatchError{
			Code:  "invalid_request",
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, types.BatchErrorResponse{Status: "error", Id: "batch_12345", Errors: types.JsonArray[types.BatchError]{batchError}})
		return
	}

	c.JSON(http.StatusAccepted, types.BatchResponse{Status: "in_progress", Id: "batch_12345"})
}

// GetSamplesBatch godoc
// @Summary Get status of a batch of samples
// @Description Get the current status of a batch job by ID
// @Tags Samples
// @Produce json
// @Param id path string true "batch id"
// @Success 201 {object} types.BatchResponse
// @Failure 400 {object} types.BatchErrorResponse
// @Failure 401
// @Security bearerauth
// @Router /samples/batch/{id} [get]
func GetSamplesBatch(c *gin.Context) {
	id := c.Param("id")
	if strings.HasPrefix(id, "error-") {
		batchErrors := types.JsonArray[types.BatchError]{
			types.BatchError{Field: "sample[0]", Error: "Patient with organization_patient_id 123 and organization_code CHOP does not exist", Code: "PATIENT_NOT_FOUND"},
			types.BatchError{Field: "sample[1]", Error: "Parent sample with sumbitter id 1234 dos not exist", Code: "PARENT_SAMPLE_NOT_FOUND"},
		}
		c.JSON(http.StatusBadRequest, types.BatchErrorResponse{Status: "error", Id: "batch_12345", Errors: batchErrors})
		return
	}
	c.JSON(http.StatusCreated, types.BatchResponse{Status: "complete", Id: id})
}
