package routes

import (
	"github.com/Andrewalifb/alpha-pos-system-company-service/api/controller"
	middleware "github.com/Andrewalifb/alpha-pos-system-company-service/api/midlleware"
	"github.com/gin-gonic/gin"
)

func PosRoleRoutes(r *gin.Engine, posRoleController controller.PosRoleController) {
	routes := r.Group("/api")

	// Apply the JWT middleware to all routes in this group
	routes.Use(middleware.JWTAuthMiddleware())

	routesV1 := routes.Group("/v1/roles")
	// Create New PosRole
	routesV1.POST("/pos_role", posRoleController.HandleCreatePosRoleRequest)
	// Get PosRole by ID
	routesV1.GET("/pos_role/:id", posRoleController.HandleReadPosRoleRequest)
	// Update Existing PosRole
	routesV1.PUT("/pos_role/:id", posRoleController.HandleUpdatePosRoleRequest)
	// Delete PosRole
	routesV1.DELETE("/pos_role/:id", posRoleController.HandleDeletePosRoleRequest)
	// Get All PosRoles
	routesV1.GET("/pos_roles", posRoleController.HandleReadAllPosRolesRequest)
}
