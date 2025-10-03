package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/radiant-network/radiant-case-api/internal/types"
)

// CreatePatientsBatch godoc
// @Summary Create a batch of patients
// @Description Accepts a JSON array of Patients
// @Tags Patients
// @Accept json
// @Produce json
// @Param payload	body		types.ListPatients	true	"List Patients"
// @Success 202 {object} types.BatchResponse
// @Failure 400 {object} types.BatchErrorResponse
// @Failure 401
// @Security bearerauth
// @Router /patients/batch [post]
func CreatePatientsBatch(c *gin.Context) {
	var batch types.ListPatients
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

// GetPatientsBatch godoc
// @Summary Get status of a batch of patients
// @Description Get the current status of a batch job by ID
// @Tags Patients
// @Produce json
// @Param id path string true "batch id"
// @Success 201 {object} types.BatchResponse
// @Failure 400 {object} types.BatchErrorResponse
// @Failure 401
// @Security bearerauth
// @Router /patients/batch/{id} [get]
func GetPatientsBatch(c *gin.Context) {
	id := c.Param("id")
	if strings.HasPrefix(id, "error-") {
		batchErrors := types.JsonArray[types.BatchError]{
			types.BatchError{Field: "patient[0]", Error: "Patient with organization_patient_id 123 and organization_code CHOP already exist", Code: "PATIENT_ALREADY_EXIST"},
		}
		c.JSON(http.StatusBadRequest, types.BatchErrorResponse{Status: "error", Id: "batch_12345", Errors: batchErrors})
		return
	}
	c.JSON(http.StatusCreated, types.BatchResponse{Status: "complete", Id: id})
}
