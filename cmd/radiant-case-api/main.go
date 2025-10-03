package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/radiant-network/radiant-case-api/internal/handlers"
	"github.com/radiant-network/radiant-case-api/internal/middleware"
)

// @title Radiant Case API
// @version 0.1.0
// @description Mock API for case creation and updates (OpenAPI 3.1 intent)
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host http://localhost:8080
// @BasePath /
// @securitydefinitions.bearerauth
func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))
	v1 := r.Group("/")
	v1.Use(middleware.OAuthMiddleware())
	{
		v1.POST("/cases/batch", handlers.CreateCasesBatch)
		v1.PATCH("/cases/:id", handlers.UpdateCase)
		v1.GET("/cases/batch/:id", handlers.GetCasesBatch)
		v1.POST("/patients/batch", handlers.CreatePatientsBatch)
		v1.GET("/patients/batch/:id", handlers.GetPatientsBatch)
		v1.POST("/samples/batch", handlers.CreateSamplesBatch)
		v1.GET("/samples/batch/:id", handlers.GetSamplesBatch)
	}

	r.Run(":8080")
}
