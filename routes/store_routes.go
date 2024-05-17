package routes

import (
	"github.com/Andrewalifb/alpha-pos-system-company-service/api/controller"
	middleware "github.com/Andrewalifb/alpha-pos-system-company-service/api/midlleware"
	"github.com/gin-gonic/gin"
)

func PosStoreRoutes(r *gin.Engine, posStoreController controller.PosStoreController) {
	routes := r.Group("/api")

	// Apply the JWT middleware to all routes in this group
	routes.Use(middleware.JWTAuthMiddleware())

	routesV1 := routes.Group("/v1/stores")
	// Create New PosStore
	routesV1.POST("/pos_store", posStoreController.HandleCreatePosStoreRequest)
	// Get PosStore by ID
	routesV1.GET("/pos_store/:id", posStoreController.HandleReadPosStoreRequest)
	// Update Existing PosStore
	routesV1.PUT("/pos_store/:id", posStoreController.HandleUpdatePosStoreRequest)
	// Delete PosStore
	routesV1.DELETE("/pos_store/:id", posStoreController.HandleDeletePosStoreRequest)
	// Get All PosStores
	routesV1.GET("/pos_stores", posStoreController.HandleReadAllPosStoresRequest)
}
