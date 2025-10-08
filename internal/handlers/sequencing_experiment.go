package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/radiant-network/radiant-case-api/internal/types"
	"net/http"
)

// AddTask godoc
// @Summary Add a task to a sequencing experiment
// @Description Add a task to a sequencing experiment
// @Tags Sequencing Experiments
// @Accept json
// @Produce json
// @Security bearerauth
// @Param id path string true "sequencing experiment id"
// @Param payload body types.SequencingExperimentTask true "task"
// @Success 202 {object} types.OperationResponse
// @Failure 400 {object} types.UpdateErrorResponse
// @Failure 401
// @Router /sequencing_experiments/{id}/task [post]
func AddTask(c *gin.Context) {
	id := c.Param("id")
	var task types.SequencingExperimentTask
	if err := c.ShouldBindJSON(&task); err != nil {
		responseError := types.UpdateErrorResponse{
			Id:    id,
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, responseError)
		return
	}

	c.JSON(http.StatusAccepted, types.OperationResponse{Status: "in_progress", OperationId: "batch_12345"})

}
