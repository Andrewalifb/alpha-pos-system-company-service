package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Andrewalifb/alpha-pos-system-company-service/api/controller"
	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}

	clientPort := os.Getenv("CLIENT_PORT")
	serverPort := os.Getenv("SERVER_PORT")

	addr := fmt.Sprintf("localhost:%s", serverPort)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Initialize the services
	branchSvc := pb.NewPosStoreBranchServiceClient(conn)
	companySvc := pb.NewPosCompanyServiceClient(conn)
	roleSvc := pb.NewPosRoleServiceClient(conn)
	storeSvc := pb.NewPosStoreServiceClient(conn)
	userSvc := pb.NewPosUserServiceClient(conn)

	// Initialize the controllers
	branchCtrl := controller.NewPosStoreBranchController(branchSvc)
	companyCtrl := controller.NewPosCompanyController(companySvc)
	roleCtrl := controller.NewPosRoleController(roleSvc)
	storeCtrl := controller.NewPosStoreController(storeSvc)
	userCtrl := controller.NewPosUserController(userSvc)

	// Create a new router
	r := gin.Default()

	// Define routes
	routes.PosStoreBranchRoute(r, branchCtrl)
	routes.PosCompanyRoutes(r, companyCtrl)
	routes.PosRoleRoutes(r, roleCtrl)
	routes.PosStoreRoutes(r, storeCtrl)
	routes.PosUserRoutes(r, userCtrl)

	// Start the server
	r.Run(":" + clientPort)
}
