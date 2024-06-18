package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	middleware "github.com/Andrewalifb/alpha-pos-system-company-service/api/midlleware"
	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/entity"
	"github.com/Andrewalifb/alpha-pos-system-company-service/pkg/repository"
	"github.com/Andrewalifb/alpha-pos-system-company-service/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	if !utils.IsSuperUser(loginRole.RoleName) && !utils.IsCompanyOrBranchUser(loginRole.RoleName) {
		return nil, errors.New("users cant create new user")
	}

	// Check if username has been avaliable on database
	isUsernameExist, err := s.userRepo.IsUsernameExist(strings.ToLower(req.PosUser.Username))
	if err != nil {
		// Log the error and return a user-friendly message
		log.Printf("Error checking if username exists: %v", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	if isUsernameExist {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("%s has been available on database, username must be unique", req.PosUser.Username))
	}

	// Email Validation
	if !utils.IsValidEmail(req.PosUser.Email) {
		return nil, errors.New(fmt.Sprintf("email %s is not valid", req.PosUser.Email))
	}

	// Check if email has been avaliable on database
	isEmailExist, err := s.userRepo.IsEmailExist(strings.ToLower(req.PosUser.Email))
	if err != nil {
		// Log the error and return a user-friendly message
		log.Printf("Error checking if email exists: %v", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	if isEmailExist {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("%s has been available on database, email must be unique", req.PosUser.Email))
	}

	// Get req data role name
	reqCreatedRole, err := s.roleRepo.ReadPosRole(req.PosUser.RoleId)
	if err != nil {
		return nil, err
	}

	superUserRole := os.Getenv("SUPER_USER_ROLE")
	companyRole := os.Getenv("COMPANY_USER_ROLE")
	branchRole := os.Getenv("BRANCH_USER_ROLE")
	storeRole := os.Getenv("STORE_USER_ROLE")

	switch loginRole.RoleName {
	case superUserRole:
		if reqCreatedRole.RoleName != companyRole {
			return nil, errors.New("a Super User role can only create Company user")
		}
	case companyRole:
		if reqCreatedRole.RoleName != branchRole && reqCreatedRole.RoleName != storeRole {
			return nil, errors.New("a Company role can only create Branch and Store user")
		}
	case branchRole:
		if reqCreatedRole.RoleName != storeRole {
			return nil, errors.New("a Branch role can only create Store user")
		}
	case storeRole:
		return nil, errors.New("a Store role cannot create new user")
	default:
		return nil, errors.New("invalid role for current user")
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

	switch reqCreatedRole.RoleName {
	// Create user for company role
	case companyRole:
		if loginRole.RoleName == superUserRole {
			// Company ID should not be empty
			if req.PosUser.CompanyId == "" {
				return nil, errors.New("error created user role name company, company id could not be empty")
			} else {
				gormUser.CompanyID = utils.ParseUUID(req.PosUser.CompanyId)
			}
			// Branch ID should be empty
			gormUser.BranchID = nil
			// Store ID should be empty
			gormUser.StoreID = nil
		}
	// Create user for branch role
	case branchRole:
		if loginRole.RoleName == companyRole {
			// Set company ID from jwt payload
			gormUser.CompanyID = utils.ParseUUID(req.JwtPayload.CompanyId)

			// Set Branch ID from req body
			if req.PosUser.BranchId == "" {
				return nil, errors.New("error created user role name branch, branch id could not be empty")
			} else {
				gormUser.BranchID = utils.ParseUUID(req.PosUser.BranchId)
			}
			// Store ID should be empty
			gormUser.StoreID = nil
		}
	// Create user for store role
	case storeRole:
		// if user role that want to create store user is company role
		if loginRole.RoleName == companyRole {
			// Set Company ID from jwt payload
			gormUser.CompanyID = utils.ParseUUID(req.JwtPayload.CompanyId)
			// Set Branch ID from req body
			if req.PosUser.BranchId == "" {
				return nil, errors.New("error created user role name store, branch id could not be empty")
			} else {
				gormUser.BranchID = utils.ParseUUID(req.PosUser.BranchId)
			}
			// Set Store ID from req body
			if req.PosUser.StoreId == "" {
				return nil, errors.New("error created user role name store, store id could not be empty")
			} else {
				gormUser.StoreID = utils.ParseUUID(req.PosUser.StoreId)
			}
		}

		// if user role that want to create store user is branch role
		if loginRole.RoleName == branchRole {
			// Set Company ID from jwt payload
			gormUser.CompanyID = utils.ParseUUID(req.JwtPayload.CompanyId)
			// Set branch ID from jwt payload
			gormUser.BranchID = utils.ParseUUID(req.JwtPayload.BranchId)
			// Set Store ID from req body
			if req.PosUser.StoreId == "" {
				return nil, errors.New("error created user role name store, store id could not be empty")
			} else {
				gormUser.StoreID = utils.ParseUUID(req.PosUser.StoreId)
			}
		}
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
	// fmt.Println("LOGIN SECRET : ", secretKey)
	user, err := s.userRepo.ReadPosUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	expectedIssuer := os.Getenv("ISSUER")
	expectedAudience := os.Getenv("AUDIENCE")

	claims := &middleware.JWTPayloadWithClaims{
		JWTPayload: &pb.JWTPayload{
			Name:      user.FirstName + " " + user.LastName,
			Role:      user.RoleId,
			CompanyId: user.CompanyId,
			BranchId:  user.BranchId,
			StoreId:   user.StoreId,
			UserId:    user.UserId,
			StandardClaims: &pb.StandardClaims{
				Audience:  expectedAudience,
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

	if !utils.IsCompanyOrBranchOrStoreUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to read all users data")
	}

	paginationResult, err := s.userRepo.ReadAllPosUsers(pagination, loginRole.RoleName, req.JwtPayload)
	if err != nil {
		return nil, err
	}

	pbPosUsers := paginationResult.Records.([]*pb.PosUser)

	return &pb.ReadAllPosUsersResponse{
		PosUsers: pbPosUsers,
		Limit:    int32(pagination.Limit),
		Page:     int32(pagination.Page),
		MaxPage:  int32(paginationResult.TotalPages),
		Count:    paginationResult.TotalRecords,
	}, nil
}

func (s *PosUserServiceServer) ReadPosUser(ctx context.Context, req *pb.ReadPosUserRequest) (*pb.ReadPosUserResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsCompanyOrBranchOrStoreUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to read users data")
	}

	posUser, err := s.userRepo.ReadPosUser(req.UserId)
	if err != nil {
		return nil, err
	}

	companyRole := os.Getenv("COMPANY_USER_ROLE")
	branchRole := os.Getenv("BRANCH_USER_ROLE")
	storeRole := os.Getenv("STORE_USER_ROLE")

	if loginRole.RoleName == companyRole {
		if !utils.VerifyCompanyUserAccess(loginRole.RoleName, posUser.CompanyId, req.JwtPayload.CompanyId) {
			return nil, errors.New("company users can only retrieve users within their company")
		}
	}

	if loginRole.RoleName == branchRole {
		if !utils.VerifyBranchUserAccess(loginRole.RoleName, posUser.BranchId, req.JwtPayload.BranchId) {
			return nil, errors.New("branch users can only retrieve users within their branch")
		}
	}

	if loginRole.RoleName == storeRole {
		if !utils.VerifyStoreUserAccess(loginRole.RoleName, posUser.StoreId, req.JwtPayload.StoreId) {
			return nil, errors.New("store users can only retrieve users within their branch")
		}
	}

	posUser.PasswordHash = ""
	return &pb.ReadPosUserResponse{
		PosUser: posUser,
	}, nil
}

func (s *PosUserServiceServer) UpdatePosUser(ctx context.Context, req *pb.UpdatePosUserRequest) (*pb.UpdatePosUserResponse, error) {
	// Extract role ID from JWT payload
	jwtRoleID := req.JwtPayload.Role

	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsCompanyOrBranchOrStoreUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to update users data")
	}

	// Get the user to be updated
	posUser, err := s.userRepo.ReadPosUser(req.PosUser.UserId)
	if err != nil {
		return nil, err
	}

	companyRole := os.Getenv("COMPANY_USER_ROLE")
	branchRole := os.Getenv("BRANCH_USER_ROLE")
	storeRole := os.Getenv("STORE_USER_ROLE")

	if loginRole.RoleName == companyRole {
		if !utils.VerifyCompanyUserAccess(loginRole.RoleName, posUser.CompanyId, req.JwtPayload.CompanyId) {
			return nil, errors.New("company users can only update users within their company")
		}
	}

	if loginRole.RoleName == branchRole {
		if !utils.VerifyBranchUserAccess(loginRole.RoleName, posUser.BranchId, req.JwtPayload.BranchId) {
			return nil, errors.New("branch users can only update users within their branch")
		}
	}

	if loginRole.RoleName == storeRole {
		if !utils.VerifyStoreUserAccess(loginRole.RoleName, posUser.StoreId, req.JwtPayload.StoreId) {
			return nil, errors.New("store users can only update users within their branch")
		}
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
	// fmt.Println("ID :", req.UserId)
	// Get user login role name
	loginRole, err := s.roleRepo.ReadPosRole(jwtRoleID)
	if err != nil {
		return nil, err
	}

	if !utils.IsCompanyOrBranchUser(loginRole.RoleName) {
		return nil, errors.New("users are not allowed to delete users data")
	}

	// Get the user to be updated
	posUser, err := s.userRepo.ReadPosUser(req.UserId)
	if err != nil {
		return nil, err
	}

	companyRole := os.Getenv("COMPANY_USER_ROLE")
	branchRole := os.Getenv("BRANCH_USER_ROLE")

	if loginRole.RoleName == companyRole {
		if !utils.VerifyCompanyUserAccess(loginRole.RoleName, posUser.CompanyId, req.JwtPayload.CompanyId) {
			return nil, errors.New("company users can only delete users within their company")
		}
	}

	if loginRole.RoleName == branchRole {
		if !utils.VerifyBranchUserAccess(loginRole.RoleName, posUser.BranchId, req.JwtPayload.BranchId) {
			return nil, errors.New("branch users can only delete users within their branch")
		}
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
