package service

import (
	"context"
	"errors"
	"time"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/entity"
	"github.com/Andrewalifb/alpha-pos-system-company-service/pkg/repository"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PosStoreService interface {
	CreatePosStore(ctx context.Context, req *pb.CreatePosStoreRequest) (*pb.CreatePosStoreResponse, error)
	ReadPosStore(ctx context.Context, req *pb.ReadPosStoreRequest) (*pb.ReadPosStoreResponse, error)
	UpdatePosStore(ctx context.Context, req *pb.UpdatePosStoreRequest) (*pb.UpdatePosStoreResponse, error)
	DeletePosStore(ctx context.Context, req *pb.DeletePosStoreRequest) (*pb.DeletePosStoreResponse, error)
	ReadAllPosStores(ctx context.Context, req *pb.ReadAllPosStoresRequest) (*pb.ReadAllPosStoresResponse, error)
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

	if loginRole.RoleName == "store" {
		return nil, errors.New("store users are not allowed to create new store id")
	}

	req.PosStore.StoreId = uuid.New().String() // Generate a new UUID for the store_id

	now := timestamppb.New(time.Now())
	req.PosStore.CreatedAt = now
	req.PosStore.UpdatedAt = now

	// Convert pb.PosStore to entity.PosStore
	gormStore := &entity.PosStore{
		StoreID:   uuid.MustParse(req.PosStore.StoreId),
		StoreName: req.PosStore.StoreName,
		BranchID:  uuid.MustParse(req.JwtPayload.BranchId),
		Location:  req.PosStore.Location,
		CompanyID: uuid.MustParse(req.JwtPayload.CompanyId),
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
	// Get req data role name
	posStore, err := s.storeRepo.ReadPosStore(req.StoreId)
	if err != nil {
		return nil, err
	}

	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	// Check if the role is "store"
	if loginRole.RoleName == "store" {
		return nil, errors.New("Store users are not allowed to retrieve stores")
	}

	// Check if the role is "company" and the company IDs don't match
	if loginRole.RoleName == "company" && posStore.CompanyId != req.JwtPayload.CompanyId {
		return nil, errors.New("Company users can only retrieve stores within their company")
	}

	// Check if the role is "branch" and the branch IDs don't match
	if loginRole.RoleName == "branch" && posStore.BranchId != req.JwtPayload.BranchId {
		return nil, errors.New("Branch users can only retrieve stores within their branch")
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

	// Check if the role is "store"
	if loginRole.RoleName == "store" {
		return nil, errors.New("Store users are not allowed to update users")
	}

	// Get the store to be updated
	posStore, err := s.storeRepo.ReadPosStore(req.PosStore.StoreId)
	if err != nil {
		return nil, err
	}

	// Check if the role is "company" and the company IDs don't match
	if loginRole.RoleName == "company" && posStore.CompanyId != req.JwtPayload.CompanyId {
		return nil, errors.New("Company users can only update stores within their company")
	}

	// Check if the role is "branch" and the branch IDs don't match
	if loginRole.RoleName == "branch" && posStore.BranchId != req.JwtPayload.BranchId {
		return nil, errors.New("Branch users can only update stores within their branch")
	}

	now := timestamppb.New(time.Now())
	req.PosStore.UpdatedAt = now
	req.PosStore.UpdatedBy = req.JwtPayload.UserId

	// Convert pb.PosStore to entity.PosStore
	gormStore := &entity.PosStore{
		StoreID:   uuid.MustParse(req.PosStore.StoreId),
		StoreName: req.PosStore.StoreName,
		BranchID:  uuid.MustParse(req.PosStore.BranchId),
		Location:  req.PosStore.Location,
		CompanyID: uuid.MustParse(req.PosStore.CompanyId),
		CreatedAt: posStore.CreatedAt.AsTime(),
		CreatedBy: uuid.MustParse(posStore.CreatedBy),
		UpdatedAt: req.PosStore.UpdatedAt.AsTime(),
		UpdatedBy: uuid.MustParse(req.PosStore.UpdatedBy),
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

	// Check if the role is "store"
	if loginRole.RoleName == "store" {
		return nil, errors.New("Store users are not allowed to delete stores")
	}

	// Get the store to be deleted
	posStore, err := s.storeRepo.ReadPosStore(req.StoreId)
	if err != nil {
		return nil, err
	}

	// Check if the role is "company" and the company IDs don't match
	if loginRole.RoleName == "company" && posStore.CompanyId != req.JwtPayload.CompanyId {
		return nil, errors.New("Company users can only delete stores within their company")
	}

	// Check if the role is "branch" and the branch IDs don't match
	if loginRole.RoleName == "branch" && posStore.BranchId != req.JwtPayload.BranchId {
		return nil, errors.New("Branch users can only delete stores within their branch")
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
