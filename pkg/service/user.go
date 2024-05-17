package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	middleware "github.com/Andrewalifb/alpha-pos-system-company-service/api/midlleware"
	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/entity"
	"github.com/Andrewalifb/alpha-pos-system-company-service/pkg/repository"
	"github.com/Andrewalifb/alpha-pos-system-company-service/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"golang.org/x/crypto/bcrypt"
)

type PosUserService interface {
	CreatePosUser(ctx context.Context, req *pb.CreatePosUserRequest) (*pb.CreatePosUserResponse, error)
	ReadPosUser(ctx context.Context, req *pb.ReadPosUserRequest) (*pb.ReadPosUserResponse, error)
	UpdatePosUser(ctx context.Context, req *pb.UpdatePosUserRequest) (*pb.UpdatePosUserResponse, error)
	DeletePosUser(ctx context.Context, req *pb.DeletePosUserRequest) (*pb.DeletePosUserResponse, error)
	ReadAllPosUsers(ctx context.Context, req *pb.ReadAllPosUsersRequest) (*pb.ReadAllPosUsersResponse, error)
	Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
}

type PosUserServiceServer struct {
	pb.UnimplementedPosUserServiceServer
	userRepo repository.PosUserRepository
	roleRepo repository.PosRoleRepository
}

func NewPosUserServiceServer(userRepo repository.PosUserRepository, roleRepo repository.PosRoleRepository) *PosUserServiceServer {
	return &PosUserServiceServer{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

func (s *PosUserServiceServer) CreatePosUser(ctx context.Context, req *pb.CreatePosUserRequest) (*pb.CreatePosUserResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	// Get req data role name
	reqCreatedRole, err := s.roleRepo.ReadPosRole(req.PosUser.RoleId)
	if err != nil {
		return nil, err
	}

	// Check role ID to determine what kind of user can be created
	switch loginRole.RoleName {
	case "super user":

	case "company":
		// Company level user can create Branch and Store level users
		if reqCreatedRole.RoleName != "branch" && reqCreatedRole.RoleName != "store" {
			return nil, errors.New("A Company role can only create Branch and Store roles")
		}
	case "branch":
		// Branch level user can only create Store level users
		if reqCreatedRole.RoleName != "store" {
			return nil, errors.New("A Branch role can only create Store roles")
		}
	case "store":
		// Store level user cannot create any users
		return nil, errors.New("A Store role cannot create new users")
	default:
		return nil, errors.New("Invalid role for current user")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.PosUser.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	req.PosUser.PasswordHash = string(hashedPassword)
	req.PosUser.UserId = uuid.New().String() // Generate a new UUID for the user_id

	now := timestamppb.New(time.Now())
	req.PosUser.CreatedAt = now
	req.PosUser.UpdatedAt = now

	// Convert pb.PosUser to entity.PosUser
	gormUser := &entity.PosUser{
		UserID:       uuid.MustParse(req.PosUser.UserId),
		Username:     req.PosUser.Username,
		PasswordHash: req.PosUser.PasswordHash,
		RoleID:       uuid.MustParse(req.PosUser.RoleId),
		CompanyID:    nil,
		BranchID:     nil,
		StoreID:      nil,
		FirstName:    req.PosUser.FirstName,
		LastName:     req.PosUser.LastName,
		Email:        req.PosUser.Email,
		PhoneNumber:  req.PosUser.PhoneNumber,
		CreatedAt:    req.PosUser.CreatedAt.AsTime(),
		CreatedBy:    uuid.MustParse(req.JwtPayload.UserId),
		UpdatedAt:    req.PosUser.UpdatedAt.AsTime(),
		UpdatedBy:    uuid.MustParse(req.JwtPayload.UserId),
	}

	if req.PosUser.CompanyId != "" {
		gormUser.CompanyID = utils.ParseUUID(req.PosUser.CompanyId)
	} else {
		gormUser.CompanyID = nil
	}

	if req.PosUser.BranchId != "" {
		gormUser.BranchID = utils.ParseUUID(req.PosUser.BranchId)
	} else {
		gormUser.BranchID = nil
	}

	if req.PosUser.StoreId != "" {
		gormUser.StoreID = utils.ParseUUID(req.PosUser.StoreId)
	} else {
		gormUser.StoreID = nil
	}

	err = s.userRepo.CreatePosUser(gormUser)
	if err != nil {
		return nil, err
	}
	req.PosUser.PasswordHash = ""
	return &pb.CreatePosUserResponse{
		PosUser: req.PosUser,
	}, nil
}

func (s *PosUserServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var secretKey = os.Getenv("SECRET_KEY")
	fmt.Println("LOGIN SECRET : ", secretKey)
	user, err := s.userRepo.ReadPosUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	expectedIssuer := os.Getenv("ISSUER")
	claims := &middleware.JWTPayloadWithClaims{
		JWTPayload: &pb.JWTPayload{
			Name:      user.FirstName + " " + user.LastName,
			Role:      user.RoleId,
			CompanyId: user.CompanyId,
			BranchId:  user.BranchId,
			StoreId:   user.StoreId,
			UserId:    user.UserId,
			StandardClaims: &pb.StandardClaims{
				Audience:  "Alpha Pos System",
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
				Id:        user.CompanyId,
				IssuedAt:  time.Now().Unix(),
				Issuer:    expectedIssuer,
				NotBefore: time.Now().Unix(),
				Subject:   user.RoleId,
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		JwtToken: ss,
	}, nil
}

func (s *PosUserServiceServer) ReadAllPosUsers(ctx context.Context, req *pb.ReadAllPosUsersRequest) (*pb.ReadAllPosUsersResponse, error) {
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

	paginationResult, err := s.userRepo.ReadAllPosUsers(pagination, loginRole.RoleName, req.JwtPayload)
	if err != nil {
		return nil, err
	}

	posUsers := paginationResult.Records.([]entity.PosUser)
	pbPosUsers := make([]*pb.PosUser, len(posUsers))

	for i, posUser := range posUsers {
		pbPosUsers[i] = &pb.PosUser{
			UserId:       posUser.UserID.String(),
			Username:     posUser.Username,
			PasswordHash: posUser.PasswordHash,
			RoleId:       posUser.RoleID.String(),
			CompanyId:    posUser.CompanyID.String(),
			BranchId:     posUser.BranchID.String(),
			StoreId:      posUser.StoreID.String(),
			FirstName:    posUser.FirstName,
			LastName:     posUser.LastName,
			Email:        posUser.Email,
			PhoneNumber:  posUser.PhoneNumber,
			CreatedAt:    timestamppb.New(posUser.CreatedAt),
			CreatedBy:    posUser.CreatedBy.String(),
			UpdatedAt:    timestamppb.New(posUser.UpdatedAt),
			UpdatedBy:    posUser.UpdatedBy.String(),
		}
	}

	return &pb.ReadAllPosUsersResponse{
		PosUsers: pbPosUsers,
		Limit:    int32(pagination.Limit),
		Page:     int32(pagination.Page),
		MaxPage:  int32(paginationResult.TotalPages),
		Count:    paginationResult.TotalRecords,
	}, nil
}

func (s *PosUserServiceServer) ReadPosUser(ctx context.Context, req *pb.ReadPosUserRequest) (*pb.ReadPosUserResponse, error) {
	// Get req data role name
	posUser, err := s.userRepo.ReadPosUser(req.UserId)
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
		return nil, errors.New("Store users are not allowed to retrieve users")
	}

	// Check if the role is "company" and the company IDs don't match
	if loginRole.RoleName == "company" && posUser.CompanyId != req.JwtPayload.CompanyId {
		return nil, errors.New("Company users can only retrieve users within their company")
	}

	// Check if the role is "branch" and the branch IDs don't match
	if loginRole.RoleName == "branch" && posUser.BranchId != req.JwtPayload.BranchId {
		return nil, errors.New("Branch users can only retrieve users within their branch")
	}

	return &pb.ReadPosUserResponse{
		PosUser: posUser,
	}, nil
}

func (s *PosUserServiceServer) UpdatePosUser(ctx context.Context, req *pb.UpdatePosUserRequest) (*pb.UpdatePosUserResponse, error) {
	// Get the role name from the role ID in the JWT payload
	// Extract role ID from JWT payload
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

	// Get the user to be updated
	posUser, err := s.userRepo.ReadPosUser(req.PosUser.UserId)
	if err != nil {
		return nil, err
	}

	// Check if the role is "company" and the company IDs don't match
	if loginRole.RoleName == "company" && posUser.CompanyId != req.JwtPayload.CompanyId {
		return nil, errors.New("Company users can only update users within their company")
	}

	// Check if the role is "branch" and the branch IDs don't match
	if loginRole.RoleName == "branch" && posUser.BranchId != req.JwtPayload.BranchId {
		return nil, errors.New("Branch users can only update users within their branch")
	}

	now := timestamppb.New(time.Now())
	req.PosUser.UpdatedAt = now

	// Convert pb.PosUser to entity.PosUser
	newUserData := &entity.PosUser{
		UserID:       uuid.MustParse(posUser.UserId), // cant be updated
		Username:     posUser.Username,               // can be updated on future update
		PasswordHash: posUser.PasswordHash,           // can be updated on future update
		RoleID:       uuid.MustParse(posUser.RoleId),
		CompanyID:    utils.ParseUUID(posUser.CompanyId), // cant be updated
		BranchID:     utils.ParseUUID(posUser.BranchId),  // cant be updated
		StoreID:      utils.ParseUUID(posUser.StoreId),   // cant be updated
		FirstName:    req.PosUser.FirstName,
		LastName:     req.PosUser.LastName,
		Email:        req.PosUser.Email,
		PhoneNumber:  req.PosUser.PhoneNumber,
		CreatedAt:    posUser.CreatedAt.AsTime(),        // cant be updated
		CreatedBy:    uuid.MustParse(posUser.CreatedBy), // cant be updated
		UpdatedAt:    req.PosUser.UpdatedAt.AsTime(),
		UpdatedBy:    uuid.MustParse(req.JwtPayload.UserId),
	}

	// Update the user
	posUser, err = s.userRepo.UpdatePosUser(newUserData)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePosUserResponse{
		PosUser: posUser,
	}, nil
}

func (s *PosUserServiceServer) DeletePosUser(ctx context.Context, req *pb.DeletePosUserRequest) (*pb.DeletePosUserResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role
	fmt.Println("ID :", req.UserId)
	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}
	fmt.Println("login data :", loginRole)
	// Check if the role is "store"
	if loginRole.RoleName == "store" {
		return nil, errors.New("store users are not allowed to delete users")
	}
	fmt.Println("ID DATA :", req.UserId)
	// Get the user to be updated
	posUser, err := s.userRepo.ReadPosUser(req.UserId)
	if err != nil {
		return nil, err
	}
	fmt.Println("DELETE DATA :", posUser)
	// Check if the role is "company" and the company IDs don't match
	if loginRole.RoleName == "company" && posUser.CompanyId != req.JwtPayload.CompanyId {
		return nil, errors.New("Company users can only delete users within their company")
	}

	// Check if the role is "branch" and the branch IDs don't match
	if loginRole.RoleName == "branch" && posUser.BranchId != req.JwtPayload.BranchId {
		return nil, errors.New("Branch users can only delete users within their branch")
	}

	// Delete the user
	err = s.userRepo.DeletePosUser(req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.DeletePosUserResponse{
		Success: true,
	}, nil
}
