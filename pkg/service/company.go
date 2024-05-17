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

type PosCompanyService interface {
	CreatePosCompany(ctx context.Context, req *pb.CreatePosCompanyRequest) (*pb.CreatePosCompanyResponse, error)
	ReadPosCompany(ctx context.Context, req *pb.ReadPosCompanyRequest) (*pb.ReadPosCompanyResponse, error)
	UpdatePosCompany(ctx context.Context, req *pb.UpdatePosCompanyRequest) (*pb.UpdatePosCompanyResponse, error)
	DeletePosCompany(ctx context.Context, req *pb.DeletePosCompanyRequest) (*pb.DeletePosCompanyResponse, error)
	ReadAllPosCompanies(ctx context.Context, req *pb.ReadAllPosCompaniesRequest) (*pb.ReadAllPosCompaniesResponse, error)
}

type PosCompanyServiceServer struct {
	pb.UnimplementedPosCompanyServiceServer
	companyRepo repository.PosCompanyRepository
	roleRepo    repository.PosRoleRepository
}

func NewPosCompanyServiceServer(companyRepo repository.PosCompanyRepository, roleRepo repository.PosRoleRepository) *PosCompanyServiceServer {
	return &PosCompanyServiceServer{
		companyRepo: companyRepo,
		roleRepo:    roleRepo,
	}
}

func (s *PosCompanyServiceServer) CreatePosCompany(ctx context.Context, req *pb.CreatePosCompanyRequest) (*pb.CreatePosCompanyResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}
	fmt.Println("LOGIN ROLE :", loginRole.RoleName)
	// Check if the role is "super user"
	if loginRole.RoleName != "super user" {
		return nil, errors.New("users are not allowed to create company roles")
	}

	req.PosCompany.CompanyId = uuid.New().String() // Generate a new UUID for the company_id

	now := timestamppb.New(time.Now())
	req.PosCompany.CreatedAt = now
	req.PosCompany.UpdatedAt = now

	// Convert pb.PosCompany to entity.PosCompany
	gormCompany := &entity.PosCompany{
		CompanyID:   uuid.MustParse(req.PosCompany.CompanyId),
		CompanyName: req.PosCompany.CompanyName,
		CreatedAt:   req.PosCompany.CreatedAt.AsTime(),
		CreatedBy:   uuid.MustParse(req.JwtPayload.UserId),
		UpdatedAt:   req.PosCompany.UpdatedAt.AsTime(),
		UpdatedBy:   uuid.MustParse(req.JwtPayload.UserId),
	}

	err = s.companyRepo.CreatePosCompany(gormCompany)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePosCompanyResponse{
		PosCompany: req.PosCompany,
	}, nil
}

func (s *PosCompanyServiceServer) ReadAllPosCompanies(ctx context.Context, req *pb.ReadAllPosCompaniesRequest) (*pb.ReadAllPosCompaniesResponse, error) {
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

	paginationResult, err := s.companyRepo.ReadAllPosCompanies(pagination, loginRole.RoleName, req.JwtPayload)
	if err != nil {
		return nil, err
	}

	posCompanies := paginationResult.Records.([]entity.PosCompany)
	pbPosCompanies := make([]*pb.PosCompany, len(posCompanies))

	for i, posCompany := range posCompanies {
		pbPosCompanies[i] = &pb.PosCompany{
			CompanyId:   posCompany.CompanyID.String(),
			CompanyName: posCompany.CompanyName,
			CreatedAt:   timestamppb.New(posCompany.CreatedAt),
			CreatedBy:   posCompany.CreatedBy.String(),
			UpdatedAt:   timestamppb.New(posCompany.UpdatedAt),
			UpdatedBy:   posCompany.UpdatedBy.String(),
		}
	}

	return &pb.ReadAllPosCompaniesResponse{
		PosCompanies: pbPosCompanies,
		Limit:        int32(pagination.Limit),
		Page:         int32(pagination.Page),
		MaxPage:      int32(paginationResult.TotalPages),
		Count:        paginationResult.TotalRecords,
	}, nil
}

func (s *PosCompanyServiceServer) ReadPosCompany(ctx context.Context, req *pb.ReadPosCompanyRequest) (*pb.ReadPosCompanyResponse, error) {
	// Get req data role name
	posCompany, err := s.companyRepo.ReadPosCompany(req.CompanyId)
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

	// Check if the role is "store" or "branch"
	if loginRole.RoleName == "store" || loginRole.RoleName == "branch" {
		return nil, errors.New("users are not allowed to retrieve companies")
	}

	// Check if the role is "company" and the company IDs don't match
	if loginRole.RoleName == "company" && posCompany.CompanyId != req.JwtPayload.CompanyId {
		return nil, errors.New("Company users can only retrieve companies within their company")
	}

	return &pb.ReadPosCompanyResponse{
		PosCompany: posCompany,
	}, nil
}

func (s *PosCompanyServiceServer) UpdatePosCompany(ctx context.Context, req *pb.UpdatePosCompanyRequest) (*pb.UpdatePosCompanyResponse, error) {
	// Get the role name from the role ID in the JWT payload
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	// Check if the role is "store"
	if loginRole.RoleName != "super user" {
		return nil, errors.New("Store users are not allowed to update companies")
	}

	// Get the company to be updated
	posCompany, err := s.companyRepo.ReadPosCompany(req.PosCompany.CompanyId)
	if err != nil {
		return nil, err
	}

	now := timestamppb.New(time.Now())
	req.PosCompany.UpdatedAt = now

	newCompanyData := &entity.PosCompany{
		CompanyID:   uuid.MustParse(posCompany.CompanyId),
		CompanyName: req.PosCompany.CompanyName,
		CreatedAt:   posCompany.CreatedAt.AsTime(),
		CreatedBy:   uuid.MustParse(posCompany.CreatedBy),
		UpdatedAt:   req.PosCompany.UpdatedAt.AsTime(),
		UpdatedBy:   uuid.MustParse(req.JwtPayload.UserId),
	}
	// Update the company
	posCompany, err = s.companyRepo.UpdatePosCompany(newCompanyData)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePosCompanyResponse{
		PosCompany: posCompany,
	}, nil
}
func (s *PosCompanyServiceServer) DeletePosCompany(ctx context.Context, req *pb.DeletePosCompanyRequest) (*pb.DeletePosCompanyResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	// Check if the role is "store"
	if loginRole.RoleName != "super user" {
		return nil, errors.New("users are not allowed to delete companies")
	}

	// Delete the company
	err = s.companyRepo.DeletePosCompany(req.CompanyId)
	if err != nil {
		return nil, err
	}

	return &pb.DeletePosCompanyResponse{
		Success: true,
	}, nil
}
