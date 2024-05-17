package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/config"
	"github.com/Andrewalifb/alpha-pos-system-company-service/pkg/repository"
	"github.com/Andrewalifb/alpha-pos-system-company-service/pkg/service"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}
	// Initialize the database
	dbConfig := config.NewConfig()

	// Initialize the repositories
	branchRepo := repository.NewPosStoreBranchRepository(dbConfig.SQLDB, dbConfig.RedisDB)
	companyRepo := repository.NewPosCompanyRepository(dbConfig.SQLDB, dbConfig.RedisDB)
	roleRepo := repository.NewPosRoleRepository(dbConfig.SQLDB, dbConfig.RedisDB)
	storeRepo := repository.NewPosStoreRepository(dbConfig.SQLDB, dbConfig.RedisDB)
	userRepo := repository.NewPosUserRepository(dbConfig.SQLDB, dbConfig.RedisDB)

	// // Initialize the services
	branchSvc := service.NewPosStoreBranchServiceServer(branchRepo, roleRepo)
	companySvc := service.NewPosCompanyServiceServer(companyRepo, roleRepo)
	roleSvc := service.NewPosRoleServiceServer(roleRepo)
	storeSvc := service.NewPosStoreServiceServer(storeRepo, roleRepo)
	userSvc := service.NewPosUserServiceServer(userRepo, roleRepo)
	// userSvc := service.NewPosUserServiceServer(userRepo, roleRepo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPosStoreBranchServiceServer(s, branchSvc)
	pb.RegisterPosCompanyServiceServer(s, companySvc)
	pb.RegisterPosRoleServiceServer(s, roleSvc)
	pb.RegisterPosStoreServiceServer(s, storeSvc)
	pb.RegisterPosUserServiceServer(s, userSvc)

	log.Println("Server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
