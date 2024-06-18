package service

import (
	"context"
	"errors"
	"os"
	"time"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/entity"
	"github.com/Andrewalifb/alpha-pos-system-company-service/pkg/repository"
	"github.com/Andrewalifb/alpha-pos-system-company-service/utils"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PosStoreService interface {
	CreatePosStore(ctx context.Context, req *pb.CreatePosStoreRequest) (*pb.CreatePosStoreResponse, error)
	ReadPosStore(ctx context.Context, req *pb.ReadPosStoreRequest) (*pb.ReadPosStoreResponse, error)
	UpdatePosStore(ctx context.Context, req *pb.UpdatePosStoreRequest) (*pb.UpdatePosStoreResponse, error)
	DeletePosStore(ctx context.Context, req *pb.DeletePosStoreRequest) (*pb.DeletePosStoreResponse, error)
	ReadAllPosStores(ctx context.Context, req *pb.ReadAllPosStoresRequest) (*pb.ReadAllPosStoresResponse, error)
	GetNextReceiptID(ctx context.Context, req *pb.GetNextReceiptIDRequest) (*pb.GetNextReceiptIDResponse, error)
}

type PosStoreServiceServer struct {
	pb.UnimplementedPosStoreServiceServer
	storeRepo repository.PosStoreRepository
	roleRepo  repository.PosRoleRepository
}

func NewPosStoreServiceServer(storeRepo repository.PosStoreRepository, roleRepo repository.PosRoleRepository) *PosStoreServiceServer {
	return &PosStoreServiceServer{
		storeRepo: storeRepo,
		roleRepo:  roleRepo,
	}
}

func (s *PosStoreServiceServer) CreatePosStore(ctx context.Context, req *pb.CreatePosStoreRequest) (*pb.CreatePosStoreResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsCompanyOrBranchUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to create store")
	}

	var companyID uuid.UUID
	var branchID uuid.UUID

	companyRole := os.Getenv("COMPANY_USER_ROLE")
	branchRole := os.Getenv("BRANCH_USER_ROLE")

	switch loginRole.RoleName {
	case companyRole:
		companyID = uuid.MustParse(req.JwtPayload.CompanyId)
		if req.PosStore.BranchId == "" {
			return nil, errors.New("error created store, branch id could not be empty")
		} else {
			branchID = uuid.MustParse(req.PosStore.BranchId)
		}
	case branchRole:
		companyID = uuid.MustParse(req.JwtPayload.CompanyId)
		branchID = uuid.MustParse(req.JwtPayload.BranchId)
	default:
		return nil, errors.New("invalid role for current user")
	}

	req.PosStore.StoreId = uuid.New().String() // Generate a new UUID for the store_id

	now := timestamppb.New(time.Now())
	req.PosStore.CreatedAt = now
	req.PosStore.UpdatedAt = now

	// Convert pb.PosStore to entity.PosStore
	gormStore := &entity.PosStore{
		StoreID:   uuid.MustParse(req.PosStore.StoreId),
		StoreName: req.PosStore.StoreName,
		BranchID:  branchID,
		Location:  req.PosStore.Location,
		CompanyID: companyID,
		CreatedAt: req.PosStore.CreatedAt.AsTime(),
		CreatedBy: uuid.MustParse(req.JwtPayload.UserId),
		UpdatedAt: req.PosStore.UpdatedAt.AsTime(),
		UpdatedBy: uuid.MustParse(req.JwtPayload.UserId),
	}

	err = s.storeRepo.CreatePosStore(gormStore)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePosStoreResponse{
		PosStore: req.PosStore,
	}, nil
}

func (s *PosStoreServiceServer) ReadAllPosStores(ctx context.Context, req *pb.ReadAllPosStoresRequest) (*pb.ReadAllPosStoresResponse, error) {
	pagination := dto.Pagination{
		Limit: int(req.Limit),
		Page:  int(req.Page),
	}

	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	// Role checker
	if !utils.IsCompanyOrBranchUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to read all stores data")
	}

	paginationResult, err := s.storeRepo.ReadAllPosStores(pagination, loginRole.RoleName, req.JwtPayload)
	if err != nil {
		return nil, err
	}

	posStores := paginationResult.Records.([]entity.PosStore)
	pbPosStores := make([]*pb.PosStore, len(posStores))

	for i, posStore := range posStores {
		pbPosStores[i] = &pb.PosStore{
			StoreId:   posStore.StoreID.String(),
			StoreName: posStore.StoreName,
			BranchId:  posStore.BranchID.String(),
			Location:  posStore.Location,
			CompanyId: posStore.CompanyID.String(),
			CreatedAt: timestamppb.New(posStore.CreatedAt),
			CreatedBy: posStore.CreatedBy.String(),
			UpdatedAt: timestamppb.New(posStore.UpdatedAt),
			UpdatedBy: posStore.UpdatedBy.String(),
		}
	}

	return &pb.ReadAllPosStoresResponse{
		PosStores: pbPosStores,
		Limit:     int32(pagination.Limit),
		Page:      int32(pagination.Page),
		MaxPage:   int32(paginationResult.TotalPages),
		Count:     paginationResult.TotalRecords,
	}, nil
}

func (s *PosStoreServiceServer) ReadPosStore(ctx context.Context, req *pb.ReadPosStoreRequest) (*pb.ReadPosStoreResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsCompanyOrBranchOrStoreUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to retrieve store")
	}

	posStore, err := s.storeRepo.ReadPosStore(req.StoreId)
	if err != nil {
		return nil, err
	}

	companyRole := os.Getenv("COMPANY_USER_ROLE")
	branchRole := os.Getenv("BRANCH_USER_ROLE")
	storeRole := os.Getenv("STORE_USER_ROLE")

	if loginRole.RoleName == companyRole {
		if !utils.VerifyCompanyUserAccess(loginRole.RoleName, posStore.CompanyId, req.JwtPayload.CompanyId) {
			return nil, errors.New("company users can only retrieve stores within their company")
		}
	}

	if loginRole.RoleName == branchRole {
		if !utils.VerifyBranchUserAccess(loginRole.RoleName, posStore.BranchId, req.JwtPayload.BranchId) {
			return nil, errors.New("branch users can only retrieve stores within their branch")
		}
	}

	if loginRole.RoleName == storeRole {
		if !utils.VerifyStoreUserAccess(loginRole.RoleName, posStore.StoreId, req.JwtPayload.StoreId) {
			return nil, errors.New("store users can only retrieve stores within their branch")
		}
	}

	return &pb.ReadPosStoreResponse{
		PosStore: posStore,
	}, nil
}

func (s *PosStoreServiceServer) UpdatePosStore(ctx context.Context, req *pb.UpdatePosStoreRequest) (*pb.UpdatePosStoreResponse, error) {
	// Get the role name from the role ID in the JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsCompanyOrBranchUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to update store")
	}

	// Get the store to be updated
	posStore, err := s.storeRepo.ReadPosStore(req.PosStore.StoreId)
	if err != nil {
		return nil, err
	}

	companyRole := os.Getenv("COMPANY_USER_ROLE")
	branchRole := os.Getenv("BRANCH_USER_ROLE")

	if loginRole.RoleName == companyRole {
		if !utils.VerifyCompanyUserAccess(loginRole.RoleName, posStore.CompanyId, req.JwtPayload.CompanyId) {
			return nil, errors.New("company users can only update stores within their company")
		}
	}

	if loginRole.RoleName == branchRole {
		if !utils.VerifyBranchUserAccess(loginRole.RoleName, posStore.BranchId, req.JwtPayload.BranchId) {
			return nil, errors.New("branch users can only update stores within their branch")
		}
	}

	now := timestamppb.New(time.Now())
	req.PosStore.UpdatedAt = now
	req.PosStore.UpdatedBy = req.JwtPayload.UserId

	// Convert pb.PosStore to entity.PosStore
	gormStore := &entity.PosStore{
		StoreID:   uuid.MustParse(posStore.StoreId), // auto
		StoreName: req.PosStore.StoreName,
		BranchID:  uuid.MustParse(posStore.BranchId), // auto
		Location:  req.PosStore.Location,
		CompanyID: uuid.MustParse(posStore.CompanyId),     // auto
		CreatedAt: posStore.CreatedAt.AsTime(),            // auto
		CreatedBy: uuid.MustParse(posStore.CreatedBy),     // auto
		UpdatedAt: req.PosStore.UpdatedAt.AsTime(),        // auto
		UpdatedBy: uuid.MustParse(req.PosStore.UpdatedBy), // auto
	}

	// Update the store
	posStore, err = s.storeRepo.UpdatePosStore(gormStore)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePosStoreResponse{
		PosStore: posStore,
	}, nil
}

func (s *PosStoreServiceServer) DeletePosStore(ctx context.Context, req *pb.DeletePosStoreRequest) (*pb.DeletePosStoreResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsCompanyUser(loginRole.RoleName) && !utils.IsBranchUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to delete stores")
	}

	// Get the store to be deleted
	posStore, err := s.storeRepo.ReadPosStore(req.StoreId)
	if err != nil {
		return nil, err
	}

	companyRole := os.Getenv("COMPANY_USER_ROLE")
	branchRole := os.Getenv("BRANCH_USER_ROLE")

	if loginRole.RoleName == companyRole {
		if !utils.VerifyCompanyUserAccess(loginRole.RoleName, posStore.CompanyId, req.JwtPayload.CompanyId) {
			return nil, errors.New("company users can only delete stores within their company")
		}
	}

	if loginRole.RoleName == branchRole {
		if !utils.VerifyBranchUserAccess(loginRole.RoleName, posStore.BranchId, req.JwtPayload.BranchId) {
			return nil, errors.New("branch users can only delete stores within their branch")
		}
	}

	// Delete the store
	err = s.storeRepo.DeletePosStore(req.StoreId)
	if err != nil {
		return nil, err
	}

	return &pb.DeletePosStoreResponse{
		Success: true,
	}, nil
}

func (s *PosStoreServiceServer) GetNextReceiptID(ctx context.Context, req *pb.GetNextReceiptIDRequest) (*pb.GetNextReceiptIDResponse, error) {
	receiptID, err := s.storeRepo.GetNextReceiptID(req.StoreId)
	if err != nil {
		return nil, err
	}
	return &pb.GetNextReceiptIDResponse{
		ReceiptId: int32(receiptID),
	}, nil
}
