package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/radiant-network/radiant-case-api/internal/types"
)

// CreateCasesBatch godoc
// @Summary Create a batch of cases
// @Description Accepts a JSON array of CaseBatch and returns created IDs
// @Tags Cases
// @Accept json
// @Produce json
// @Param payload	body		types.ListCaseBatches	true	"List Cases"
// @Success 202 {object} types.BatchResponse
// @Failure 400 {object} types.BatchErrorResponse
// @Failure 401
// @Security bearerauth
// @Router /cases/batch [post]
func CreateCasesBatch(c *gin.Context) {
	var batch types.ListCaseBatches
	if err := c.ShouldBindJSON(&batch); err != nil {
		batchError := types.BatchError{
			Code:  "invalid_request",
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, types.BatchErrorResponse{Status: "error", Id: "batch_12345", Errors: types.JsonArray[types.BatchError]{batchError}})
		return
	}
	cases := batch.Cases
	ids := make([]string, 0, len(cases))
	for i := range cases {
		ids = append(ids, "case-"+cases[i].ProjectCode)
	}

	c.JSON(http.StatusAccepted, types.BatchResponse{Status: "in_progress", Id: "batch_12345"})
}

// GetCasesBatch godoc
// @Summary Get status of a batch of cases
// @Description Get the current status of a batch job by ID
// @Tags Cases
// @Produce json
// @Param id path string true "batch id"
// @Success 201 {object} types.BatchResponse
// @Failure 400 {object} types.BatchErrorResponse
// @Failure 401
// @Security bearerauth
// @Router /cases/batch/{id} [get]
func GetCasesBatch(c *gin.Context) {
	id := c.Param("id")
	if strings.HasPrefix(id, "error-") {
		batchErrors := types.JsonArray[types.BatchError]{
			types.BatchError{Field: "case[0].patient[0]", Error: "Patient with organization_patient_id 123 and organization_code CHOP not found", Code: "PATIENT_NOT_FOUND"},
			types.BatchError{Field: "case[0].sequencing_experiment[0]", Error: "Sample with submitter_sample_id 123 not found", Code: "SAMPLE_NOT_FOUND"},
			types.BatchError{Field: "case[0].sequencing_experiment[1].task[3]", Error: "Task type with code unknown not supported", Code: "BAD_TASK_TYPE"},
		}
		c.JSON(http.StatusBadRequest, types.BatchErrorResponse{Status: "error", Id: "batch_12345", Errors: batchErrors})
		return
	}
	c.JSON(http.StatusCreated, types.BatchResponse{Status: "complete", Id: id})
}

// UpdateCase godoc
// @Summary Update one case
// @Description Update a case partial fields by id
// @Tags Cases
// @Accept json
// @Produce json
// @Security bearerauth
// @Param id path string true "case id"
// @Param payload body types.PartialCase true "partial fields"
// @Success 200 {object} types.PartialCaseResponse
// @Failure 400 {object} types.PartialCaseErrorResponse
// @Failure 401
// @Router /cases/{id} [patch]
func UpdateCase(c *gin.Context) {
	id := c.Param("id")
	var patch types.PartialCase
	if err := c.ShouldBindJSON(&patch); err != nil {
		responseError := types.PartialCaseErrorResponse{
			Id:    id,
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	c.JSON(http.StatusOK, types.PartialCaseResponse{Id: id})
}
