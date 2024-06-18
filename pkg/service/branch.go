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

type PosStoreBranchService interface {
	CreatePosStoreBranch(ctx context.Context, req *pb.CreatePosStoreBranchRequest) (*pb.CreatePosStoreBranchResponse, error)
	ReadPosStoreBranch(ctx context.Context, req *pb.ReadPosStoreBranchRequest) (*pb.ReadPosStoreBranchResponse, error)
	UpdatePosStoreBranch(ctx context.Context, req *pb.UpdatePosStoreBranchRequest) (*pb.UpdatePosStoreBranchResponse, error)
	DeletePosStoreBranch(ctx context.Context, req *pb.DeletePosStoreBranchRequest) (*pb.DeletePosStoreBranchResponse, error)
	ReadAllPosStoreBranches(ctx context.Context, req *pb.ReadAllPosStoreBranchesRequest) (*pb.ReadAllPosStoreBranchesResponse, error)
}

type PosStoreBranchServiceServer struct {
	pb.UnimplementedPosStoreBranchServiceServer
	branchRepo repository.PosStoreBranchRepository
	roleRepo   repository.PosRoleRepository
}

func NewPosStoreBranchServiceServer(branchRepo repository.PosStoreBranchRepository, roleRepo repository.PosRoleRepository) *PosStoreBranchServiceServer {
	return &PosStoreBranchServiceServer{
		branchRepo: branchRepo,
		roleRepo:   roleRepo,
	}
}

func (s *PosStoreBranchServiceServer) CreatePosStoreBranch(ctx context.Context, req *pb.CreatePosStoreBranchRequest) (*pb.CreatePosStoreBranchResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	// Role checker
	if !utils.IsCompanyUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to create branch")
	}

	req.PosStoreBranch.BranchId = uuid.New().String() // Generate a new UUID for the branch_id

	now := timestamppb.New(time.Now())
	req.PosStoreBranch.CreatedAt = now
	req.PosStoreBranch.UpdatedAt = now

	// Convert pb.PosStoreBranch to entity.PosStoreBranch
	gormStoreBranch := &entity.PosStoreBranch{
		BranchID:   uuid.MustParse(req.PosStoreBranch.BranchId), // auto
		BranchName: req.PosStoreBranch.BranchName,
		CompanyID:  uuid.MustParse(req.JwtPayload.CompanyId), // auto
		CreatedAt:  req.PosStoreBranch.CreatedAt.AsTime(),    // auto
		CreatedBy:  uuid.MustParse(req.JwtPayload.UserId),    // auto
		UpdatedAt:  req.PosStoreBranch.UpdatedAt.AsTime(),    // auto
		UpdatedBy:  uuid.MustParse(req.JwtPayload.UserId),    // auto
	}

	err = s.branchRepo.CreatePosStoreBranch(gormStoreBranch)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePosStoreBranchResponse{
		PosStoreBranch: req.PosStoreBranch,
	}, nil
}

func (s *PosStoreBranchServiceServer) ReadAllPosStoreBranches(ctx context.Context, req *pb.ReadAllPosStoreBranchesRequest) (*pb.ReadAllPosStoreBranchesResponse, error) {
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
	if !utils.IsCompanyUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to read all branch data")
	}

	paginationResult, err := s.branchRepo.ReadAllPosStoreBranches(pagination, loginRole.RoleName, req.JwtPayload)
	if err != nil {
		return nil, err
	}

	posStoreBranches := paginationResult.Records.([]entity.PosStoreBranch)
	pbPosStoreBranches := make([]*pb.PosStoreBranch, len(posStoreBranches))

	for i, posStoreBranch := range posStoreBranches {
		pbPosStoreBranches[i] = &pb.PosStoreBranch{
			BranchId:   posStoreBranch.BranchID.String(),
			BranchName: posStoreBranch.BranchName,
			CompanyId:  posStoreBranch.CompanyID.String(),
			CreatedAt:  timestamppb.New(posStoreBranch.CreatedAt),
			CreatedBy:  posStoreBranch.CreatedBy.String(),
			UpdatedAt:  timestamppb.New(posStoreBranch.UpdatedAt),
			UpdatedBy:  posStoreBranch.UpdatedBy.String(),
		}
	}

	return &pb.ReadAllPosStoreBranchesResponse{
		PosStoreBranches: pbPosStoreBranches,
		Limit:            int32(pagination.Limit),
		Page:             int32(pagination.Page),
		MaxPage:          int32(paginationResult.TotalPages),
		Count:            paginationResult.TotalRecords,
	}, nil
}

func (s *PosStoreBranchServiceServer) ReadPosStoreBranch(ctx context.Context, req *pb.ReadPosStoreBranchRequest) (*pb.ReadPosStoreBranchResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsCompanyOrBranchUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to retrieve store branches")
	}

	posStoreBranch, err := s.branchRepo.ReadPosStoreBranch(req.BranchId)
	if err != nil {
		return nil, err
	}

	companyRole := os.Getenv("COMPANY_USER_ROLE")
	branchRole := os.Getenv("BRANCH_USER_ROLE")

	if loginRole.RoleName == companyRole {
		if !utils.VerifyCompanyUserAccess(loginRole.RoleName, posStoreBranch.CompanyId, req.JwtPayload.CompanyId) {
			return nil, errors.New("company users can only retrieve branches within their company")
		}
	}

	if loginRole.RoleName == branchRole {
		if !utils.VerifyBranchUserAccess(loginRole.RoleName, posStoreBranch.BranchId, req.JwtPayload.BranchId) {
			return nil, errors.New("branch users can only retrieve their branch")
		}
	}

	return &pb.ReadPosStoreBranchResponse{
		PosStoreBranch: posStoreBranch,
	}, nil
}

func (s *PosStoreBranchServiceServer) UpdatePosStoreBranch(ctx context.Context, req *pb.UpdatePosStoreBranchRequest) (*pb.UpdatePosStoreBranchResponse, error) {
	// Get the role name from the role ID in the JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	// Role checker
	if !utils.IsCompanyUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to update store branches")
	}

	// Get the store branch to be updated
	posStoreBranch, err := s.branchRepo.ReadPosStoreBranch(req.PosStoreBranch.BranchId)
	if err != nil {
		return nil, err
	}

	if !utils.VerifyCompanyUserAccess(loginRole.RoleName, posStoreBranch.CompanyId, req.JwtPayload.CompanyId) {
		return nil, errors.New("company users can only update store branches within their company")
	}

	now := timestamppb.New(time.Now())
	req.PosStoreBranch.UpdatedAt = now
	req.PosStoreBranch.UpdatedBy = req.JwtPayload.UserId

	// Convert pb.PosStoreBranch to entity.PosStoreBranch
	gormStoreBranch := &entity.PosStoreBranch{
		BranchID:   uuid.MustParse(posStoreBranch.BranchId), // auto
		BranchName: req.PosStoreBranch.BranchName,
		CompanyID:  uuid.MustParse(posStoreBranch.CompanyId),     // auto
		CreatedAt:  posStoreBranch.CreatedAt.AsTime(),            // auto
		CreatedBy:  uuid.MustParse(posStoreBranch.CreatedBy),     // auto
		UpdatedAt:  req.PosStoreBranch.UpdatedAt.AsTime(),        // auto
		UpdatedBy:  uuid.MustParse(req.PosStoreBranch.UpdatedBy), // auto
	}

	// Update the store branch
	posStoreBranch, err = s.branchRepo.UpdatePosStoreBranch(gormStoreBranch)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePosStoreBranchResponse{
		PosStoreBranch: posStoreBranch,
	}, nil
}

func (s *PosStoreBranchServiceServer) DeletePosStoreBranch(ctx context.Context, req *pb.DeletePosStoreBranchRequest) (*pb.DeletePosStoreBranchResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	// Role checker
	if !utils.IsCompanyUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to delete store branches")
	}

	// Get the store branch to be deleted
	posStoreBranch, err := s.branchRepo.ReadPosStoreBranch(req.BranchId)
	if err != nil {
		return nil, err
	}

	if !utils.VerifyCompanyUserAccess(loginRole.RoleName, posStoreBranch.CompanyId, req.JwtPayload.CompanyId) {
		return nil, errors.New("company users can only delete store branches within their company")
	}

	// Delete the store branch
	err = s.branchRepo.DeletePosStoreBranch(req.BranchId)
	if err != nil {
		return nil, err
	}

	return &pb.DeletePosStoreBranchResponse{
		Success: true,
	}, nil
}
