package routes

import (
	"github.com/Andrewalifb/alpha-pos-system-company-service/api/controller"
	middleware "github.com/Andrewalifb/alpha-pos-system-company-service/api/midlleware"
	"github.com/gin-gonic/gin"
)

func PosStoreBranchRoute(r *gin.Engine, posStoreBranchController controller.PosStoreBranchController) {
	routes := r.Group("/api")

	// Apply the JWT middleware to all routes in this group
	routes.Use(middleware.JWTAuthMiddleware())

	routesV1 := routes.Group("/v1")
	// Create New PosStoreBranch
	routesV1.POST("/pos_store_branch", posStoreBranchController.HandleCreatePosStoreBranchRequest)
	// Get PosStoreBranch by ID
	routesV1.GET("/pos_store_branch/:id", posStoreBranchController.HandleReadPosStoreBranchRequest)
	// Update Existing PosStoreBranch
	routesV1.PUT("/pos_store_branch/:id", posStoreBranchController.HandleUpdatePosStoreBranchRequest)
	// Delete PosStoreBranch
	routesV1.DELETE("/pos_store_branch/:id", posStoreBranchController.HandleDeletePosStoreBranchRequest)
	// Get All PosStoreBranches
	routesV1.GET("/pos_store_branches", posStoreBranchController.HandleReadAllPosStoreBranchesRequest)
}
