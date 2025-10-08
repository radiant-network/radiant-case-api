package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/radiant-network/radiant-case-api/internal/types"
	"net/http"
)

// CreateTumorNormalTask godoc
// @Summary Create a tumor normal task
// @Description Create a tumor normal task
// @Tags Task
// @Accept json
// @Produce json
// @Security bearerauth
// @Param payload body types.TumorNormalTask true "task"
// @Success 202 {object} types.OperationResponse
// @Failure 400 {object} types.UpdateErrorResponse
// @Failure 401
// @Router /tumor_normal_tasks [post]
func CreateTumorNormalTask(c *gin.Context) {
	id := c.Param("id")
	var task types.TumorNormalTask
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
