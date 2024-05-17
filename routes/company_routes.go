package routes

import (
	"github.com/Andrewalifb/alpha-pos-system-company-service/api/controller"
	middleware "github.com/Andrewalifb/alpha-pos-system-company-service/api/midlleware"
	"github.com/gin-gonic/gin"
)

func PosCompanyRoutes(r *gin.Engine, posCompanyController controller.PosCompanyController) {
	routes := r.Group("/api")

	// Apply the JWT middleware to all routes in this group
	routes.Use(middleware.JWTAuthMiddleware())

	routesV1 := routes.Group("/v1/companies")
	// Create New PosCompany
	routesV1.POST("/pos_company", posCompanyController.HandleCreatePosCompanyRequest)
	// Get PosCompany by ID
	routesV1.GET("/pos_company/:id", posCompanyController.HandleReadPosCompanyRequest)
	// Update Existing PosCompany
	routesV1.PUT("/pos_company/:id", posCompanyController.HandleUpdatePosCompanyRequest)
	// Delete PosCompany
	routesV1.DELETE("/pos_company/:id", posCompanyController.HandleDeletePosCompanyRequest)
	// Get All PosCompanies
	routesV1.GET("/pos_companies", posCompanyController.HandleReadAllPosCompaniesRequest)
}
