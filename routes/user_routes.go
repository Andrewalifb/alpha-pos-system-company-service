package routes

import (
	"github.com/Andrewalifb/alpha-pos-system-company-service/api/controller"
	middleware "github.com/Andrewalifb/alpha-pos-system-company-service/api/midlleware"
	"github.com/gin-gonic/gin"
)

func PosUserRoutes(r *gin.Engine, posUserController controller.PosUserController) {
	routes := r.Group("/api")

	routesV1 := routes.Group("/v1/users")
	// Login route is not protected by the JWT middleware
	routesV1.POST("/login", posUserController.HandleLoginRequest)

	// Apply the JWT middleware to all other routes in this group
	routesV1.Use(middleware.JWTAuthMiddleware())

	// Create New PosUser
	routesV1.POST("/pos_user", posUserController.HandleCreatePosUserRequest)
	// Get PosUser by ID
	routesV1.GET("/pos_user/:id", posUserController.HandleReadPosUserRequest)
	// Update Existing PosUser
	routesV1.PUT("/pos_user/:id", posUserController.HandleUpdatePosUserRequest)
	// Delete PosUser
	routesV1.DELETE("/pos_user/:id", posUserController.HandleDeletePosUserRequest)
	// Get All PosUsers
	routesV1.GET("/pos_users", posUserController.HandleReadAllPosUsersRequest)
}
