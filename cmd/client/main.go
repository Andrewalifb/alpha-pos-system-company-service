package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Andrewalifb/alpha-pos-system-company-service/api/controller"
	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	// Define your routes
	routes.PosStoreBranchRoute(r, branchCtrl)
	routes.PosCompanyRoutes(r, companyCtrl)
	routes.PosRoleRoutes(r, roleCtrl)
	routes.PosStoreRoutes(r, storeCtrl)
	routes.PosUserRoutes(r, userCtrl)

	// Start the server
	r.Run(":8080")
}
