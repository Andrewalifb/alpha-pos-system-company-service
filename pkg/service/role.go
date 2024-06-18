package service

import (
	"context"
	"errors"
	"time"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/entity"
	"github.com/Andrewalifb/alpha-pos-system-company-service/pkg/repository"
	"github.com/Andrewalifb/alpha-pos-system-company-service/utils"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PosRoleService interface {
	CreatePosRole(ctx context.Context, req *pb.CreatePosRoleRequest) (*pb.CreatePosRoleResponse, error)
	ReadPosRole(ctx context.Context, req *pb.ReadPosRoleRequest) (*pb.ReadPosRoleResponse, error)
	UpdatePosRole(ctx context.Context, req *pb.UpdatePosRoleRequest) (*pb.UpdatePosRoleResponse, error)
	DeletePosRole(ctx context.Context, req *pb.DeletePosRoleRequest) (*pb.DeletePosRoleResponse, error)
	ReadAllPosRoles(ctx context.Context, req *pb.ReadAllPosRolesRequest) (*pb.ReadAllPosRolesResponse, error)
}

type PosRoleServiceServer struct {
	pb.UnimplementedPosRoleServiceServer
	repo repository.PosRoleRepository
}

func NewPosRoleServiceServer(repo repository.PosRoleRepository) *PosRoleServiceServer {
	return &PosRoleServiceServer{
		repo: repo,
	}
}

func (s *PosRoleServiceServer) CreatePosRole(ctx context.Context, req *pb.CreatePosRoleRequest) (*pb.CreatePosRoleResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.repo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsSuperUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to create roles")
	}

	req.PosRole.RoleId = uuid.New().String() // Generate a new UUID for the role_id

	now := timestamppb.New(time.Now())
	req.PosRole.CreatedAt = now
	req.PosRole.UpdatedAt = now

	// Convert pb.PosRole to entity.PosRole
	gormRole := &entity.PosRole{
		RoleID:    uuid.MustParse(req.PosRole.RoleId),
		RoleName:  req.PosRole.RoleName,
		CreatedAt: req.PosRole.CreatedAt.AsTime(),
		CreatedBy: uuid.MustParse(req.JwtPayload.UserId),
		UpdatedAt: req.PosRole.UpdatedAt.AsTime(),
		UpdatedBy: uuid.MustParse(req.JwtPayload.UserId),
	}

	err = s.repo.CreatePosRole(gormRole)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePosRoleResponse{
		PosRole: req.PosRole,
	}, nil
}

func (s *PosRoleServiceServer) ReadAllPosRoles(ctx context.Context, req *pb.ReadAllPosRolesRequest) (*pb.ReadAllPosRolesResponse, error) {
	pagination := dto.Pagination{
		Limit: int(req.Limit),
		Page:  int(req.Page),
	}

	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.repo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsSuperUserOrCompany(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to retrieve all roles")
	}

	paginationResult, err := s.repo.ReadAllPosRoles(pagination, loginRole.RoleName, req.JwtPayload)
	if err != nil {
		return nil, err
	}

	posRoles := paginationResult.Records.([]entity.PosRole)
	pbPosRoles := make([]*pb.PosRole, len(posRoles))

	for i, posRole := range posRoles {
		pbPosRoles[i] = &pb.PosRole{
			RoleId:    posRole.RoleID.String(),
			RoleName:  posRole.RoleName,
			CreatedAt: timestamppb.New(posRole.CreatedAt),
			CreatedBy: posRole.CreatedBy.String(),
			UpdatedAt: timestamppb.New(posRole.UpdatedAt),
			UpdatedBy: posRole.UpdatedBy.String(),
		}
	}

	return &pb.ReadAllPosRolesResponse{
		PosRoles: pbPosRoles,
		Limit:    int32(pagination.Limit),
		Page:     int32(pagination.Page),
		MaxPage:  int32(paginationResult.TotalPages),
		Count:    paginationResult.TotalRecords,
	}, nil
}

func (s *PosRoleServiceServer) ReadPosRole(ctx context.Context, req *pb.ReadPosRoleRequest) (*pb.ReadPosRoleResponse, error) {
	// Get req data role name
	posRole, err := s.repo.ReadPosRole(req.RoleId)
	if err != nil {
		return nil, err
	}

	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.repo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsSuperUser(loginRole.RoleName) && !utils.IsCompanyOrBranchOrStoreUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to retrieve roles")
	}

	// Convert entity.PosRole to pb.PosRole
	pbPosRole := &pb.PosRole{
		RoleId:    posRole.RoleId,
		RoleName:  posRole.RoleName,
		CreatedAt: timestamppb.New(posRole.CreatedAt.AsTime()),
		CreatedBy: posRole.CreatedBy,
		UpdatedAt: timestamppb.New(posRole.UpdatedAt.AsTime()),
		UpdatedBy: posRole.UpdatedBy,
	}

	return &pb.ReadPosRoleResponse{
		PosRole: pbPosRole,
	}, nil
}

func (s *PosRoleServiceServer) UpdatePosRole(ctx context.Context, req *pb.UpdatePosRoleRequest) (*pb.UpdatePosRoleResponse, error) {
	// Get the role name from the role ID in the JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.repo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsSuperUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to update roles")
	}

	// Get the role to be updated
	posRole, err := s.repo.ReadPosRole(req.PosRole.RoleId)
	if err != nil {
		return nil, err
	}

	now := timestamppb.New(time.Now())
	req.PosRole.UpdatedAt = now
	req.PosRole.UpdatedBy = req.JwtPayload.UserId

	// Convert pb.PosRole to entity.PosRole
	gormRole := &entity.PosRole{
		RoleID:    uuid.MustParse(posRole.RoleId), // auto
		RoleName:  req.PosRole.RoleName,
		CreatedAt: posRole.CreatedAt.AsTime(),            // auto
		CreatedBy: uuid.MustParse(posRole.CreatedBy),     // auto
		UpdatedAt: req.PosRole.UpdatedAt.AsTime(),        // auto
		UpdatedBy: uuid.MustParse(req.PosRole.UpdatedBy), // auto
	}

	// Update the role
	posRole, err = s.repo.UpdatePosRole(gormRole)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePosRoleResponse{
		PosRole: posRole,
	}, nil
}

func (s *PosRoleServiceServer) DeletePosRole(ctx context.Context, req *pb.DeletePosRoleRequest) (*pb.DeletePosRoleResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.repo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsSuperUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to delete roles")
	}

	// Delete the role
	err = s.repo.DeletePosRole(req.RoleId)
	if err != nil {
		return nil, err
	}

	return &pb.DeletePosRoleResponse{
		Success: true,
	}, nil
}
