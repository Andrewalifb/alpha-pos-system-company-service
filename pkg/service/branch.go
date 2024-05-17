package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/entity"
	"github.com/Andrewalifb/alpha-pos-system-company-service/pkg/repository"

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
	fmt.Println("ROLE NAME :", loginRole.RoleName)
	// Check if the role is "super user"
	if loginRole.RoleName != "company" {
		return nil, errors.New("users are not allowed to create branch")
	}

	// Check role ID to determine what kind of user can be created
	// switch loginRole.RoleName {
	// case "super user":
	// 	// Can create users with role "Company", "Branch", and "Store"
	// 	if req.PosStoreBranch.RoleName != "Company" && req.PosStoreBranch.RoleName != "Branch" && req.PosStoreBranch.RoleName != "Store" {
	// 		return nil, errors.New("Invalid role for new user")
	// 	}
	// case "company":
	// 	// Can create users with role "Branch" and "Store"
	// 	if req.PosStoreBranch.RoleName != "Branch" && req.PosStoreBranch.RoleName != "Store" {
	// 		return nil, errors.New("Invalid role for new user")
	// 	}
	// case "branch":
	// 	// Can create users with role "Store"
	// 	if req.PosStoreBranch.RoleName != "Store" {
	// 		return nil, errors.New("Invalid role for new user")
	// 	}
	// case "store":
	// 	// Cannot create any user
	// 	return nil, errors.New("Cannot create new user")
	// default:
	// 	return nil, errors.New("Invalid role for current user")
	// }

	req.PosStoreBranch.BranchId = uuid.New().String() // Generate a new UUID for the branch_id

	now := timestamppb.New(time.Now())
	req.PosStoreBranch.CreatedAt = now
	req.PosStoreBranch.UpdatedAt = now

	// Convert pb.PosStoreBranch to entity.PosStoreBranch
	gormStoreBranch := &entity.PosStoreBranch{
		BranchID:   uuid.MustParse(req.PosStoreBranch.BranchId),
		BranchName: req.PosStoreBranch.BranchName,
		CompanyID:  uuid.MustParse(req.JwtPayload.CompanyId),
		CreatedAt:  req.PosStoreBranch.CreatedAt.AsTime(),
		CreatedBy:  uuid.MustParse(req.JwtPayload.UserId),
		UpdatedAt:  req.PosStoreBranch.UpdatedAt.AsTime(),
		UpdatedBy:  uuid.MustParse(req.JwtPayload.UserId),
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
	// Get req data role name
	posStoreBranch, err := s.branchRepo.ReadPosStoreBranch(req.BranchId)
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
		return nil, errors.New("store users are not allowed to retrieve store branches")
	}

	// Check if the role is "company" and the company IDs don't match
	if loginRole.RoleName == "company" && posStoreBranch.CompanyId != req.JwtPayload.CompanyId {
		return nil, errors.New("company users can only retrieve store branches within their company")
	}

	// Check if the role is "branch" and the branch IDs don't match
	if loginRole.RoleName == "branch" && posStoreBranch.BranchId != req.JwtPayload.BranchId {
		return nil, errors.New("branch users can only retrieve store branches within their branch")
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

	// Check if the role is "store"
	if loginRole.RoleName == "store" {
		return nil, errors.New("store users are not allowed to update store branches")
	}

	// Get the store branch to be updated
	posStoreBranch, err := s.branchRepo.ReadPosStoreBranch(req.PosStoreBranch.BranchId)
	if err != nil {
		return nil, err
	}

	// Check if the role is "company" and the company IDs don't match
	if loginRole.RoleName == "company" && posStoreBranch.CompanyId != req.JwtPayload.CompanyId {
		return nil, errors.New("company users can only update store branches within their company")
	}

	// Check if the role is "branch" and the branch IDs don't match
	if loginRole.RoleName == "branch" && posStoreBranch.BranchId != req.JwtPayload.BranchId {
		return nil, errors.New("branch users can only update store branches within their branch")
	}

	now := timestamppb.New(time.Now())
	req.PosStoreBranch.UpdatedAt = now
	req.PosStoreBranch.UpdatedBy = req.JwtPayload.UserId

	// Convert pb.PosStoreBranch to entity.PosStoreBranch
	gormStoreBranch := &entity.PosStoreBranch{
		BranchID:   uuid.MustParse(req.PosStoreBranch.BranchId),
		BranchName: req.PosStoreBranch.BranchName,
		CompanyID:  uuid.MustParse(req.PosStoreBranch.CompanyId),
		CreatedAt:  posStoreBranch.CreatedAt.AsTime(),
		CreatedBy:  uuid.MustParse(posStoreBranch.CreatedBy),
		UpdatedAt:  req.PosStoreBranch.UpdatedAt.AsTime(),
		UpdatedBy:  uuid.MustParse(req.PosStoreBranch.UpdatedBy),
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

	// Check if the role is "store"
	if loginRole.RoleName == "store" || loginRole.RoleName == "branch" {
		return nil, errors.New("store users are not allowed to delete store branches")
	}

	// Get the store branch to be deleted
	posStoreBranch, err := s.branchRepo.ReadPosStoreBranch(req.BranchId)
	if err != nil {
		return nil, err
	}

	// Check if the role is "company" and the company IDs don't match
	if loginRole.RoleName == "company" && posStoreBranch.CompanyId != req.JwtPayload.CompanyId {
		return nil, errors.New("Company users can only delete store branches within their company")
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
