package repository

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"os"
	"time"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/entity"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PosUserRepository interface {
	CreatePosUser(posUser *entity.PosUser) error
	ReadPosUser(userID string) (*pb.PosUser, error)
	ReadPosUserByUsername(username string) (*pb.PosUser, error)
	IsUsernameExist(username string) (bool, error)
	IsEmailExist(email string) (bool, error)
	UpdatePosUser(posUser *entity.PosUser) (*pb.PosUser, error)
	DeletePosUser(userID string) error
	ReadAllPosUsers(pagination dto.Pagination, roleName string, jwtPayload *pb.JWTPayload) (*dto.PaginationResult, error)
}

type posUserRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosUserRepository(db *gorm.DB, redis *redis.Client) PosUserRepository {
	return &posUserRepository{
		db:    db,
		redis: redis,
	}
}

func (r *posUserRepository) CreatePosUser(posUser *entity.PosUser) error {
	result := r.db.Create(posUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *posUserRepository) ReadAllPosUsers(pagination dto.Pagination, roleName string, jwtPayload *pb.JWTPayload) (*dto.PaginationResult, error) {
	var posUsers []entity.PosUser
	var totalRecords int64

	query := r.db.Model(&entity.PosUser{})

	companyRole := os.Getenv("COMPANY_USER_ROLE")
	branchRole := os.Getenv("BRANCH_USER_ROLE")
	storeRole := os.Getenv("STORE_USER_ROLE")

	switch roleName {
	case companyRole:
		query = query.Where("company_id = ?", jwtPayload.CompanyId)
	case branchRole:
		query = query.Where("branch_id = ?", jwtPayload.BranchId)
	case storeRole:
		query = query.Where("store_id = ?", jwtPayload.StoreId)
	default:
		return nil, errors.New("invalid role")
	}

	if pagination.Limit > 0 && pagination.Page > 0 {
		offset := (pagination.Page - 1) * pagination.Limit
		query = query.Offset(offset).Limit(pagination.Limit)
	}

	query.Find(&posUsers)
	query.Count(&totalRecords)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pagination.Limit)))

	posUsersProto := make([]*pb.PosUser, len(posUsers))

	for i, posUserEntity := range posUsers {
		companyId := ""
		if posUserEntity.CompanyID != nil {
			companyId = posUserEntity.CompanyID.String()
		}

		branchId := ""
		if posUserEntity.BranchID != nil {
			branchId = posUserEntity.BranchID.String()
		}

		storeId := ""
		if posUserEntity.StoreID != nil {
			storeId = posUserEntity.StoreID.String()
		}

		posUsersProto[i] = &pb.PosUser{
			UserId:       posUserEntity.UserID.String(),
			Username:     posUserEntity.Username,
			PasswordHash: posUserEntity.PasswordHash,
			RoleId:       posUserEntity.RoleID.String(),
			CompanyId:    companyId,
			BranchId:     branchId,
			StoreId:      storeId,
			FirstName:    posUserEntity.FirstName,
			LastName:     posUserEntity.LastName,
			Email:        posUserEntity.Email,
			PhoneNumber:  posUserEntity.PhoneNumber,
			CreatedAt:    timestamppb.New(posUserEntity.CreatedAt),
			CreatedBy:    posUserEntity.CreatedBy.String(),
			UpdatedAt:    timestamppb.New(posUserEntity.UpdatedAt),
			UpdatedBy:    posUserEntity.UpdatedBy.String(),
		}
	}

	return &dto.PaginationResult{
		TotalRecords: totalRecords,
		Records:      posUsersProto,
		CurrentPage:  pagination.Page,
		TotalPages:   totalPages,
	}, nil
}

func (r *posUserRepository) ReadPosUser(userID string) (*pb.PosUser, error) {
	// Try to get the user from Redis first
	userData, err := r.redis.Get(context.Background(), userID).Result()

	if err == redis.Nil {

		// User not found in Redis, get from PostgreSQL
		var posUserEntity entity.PosUser
		if err := r.db.Where("user_id = ?", userID).First(&posUserEntity).Error; err != nil {
			return nil, err
		}

		// Check if the UUID fields are nil before trying to convert them
		companyId := ""
		if posUserEntity.CompanyID != nil {
			companyId = posUserEntity.CompanyID.String()
		}

		branchId := ""
		if posUserEntity.BranchID != nil {
			branchId = posUserEntity.BranchID.String()
		}

		storeId := ""
		if posUserEntity.StoreID != nil {
			storeId = posUserEntity.StoreID.String()
		}

		// Convert entity.PosUser to pb.PosUser
		protoUser := &pb.PosUser{
			UserId:       posUserEntity.UserID.String(),
			Username:     posUserEntity.Username,
			PasswordHash: posUserEntity.PasswordHash,
			RoleId:       posUserEntity.RoleID.String(),
			CompanyId:    companyId,
			BranchId:     branchId,
			StoreId:      storeId,
			FirstName:    posUserEntity.FirstName,
			LastName:     posUserEntity.LastName,
			Email:        posUserEntity.Email,
			PhoneNumber:  posUserEntity.PhoneNumber,
			CreatedAt:    timestamppb.New(posUserEntity.CreatedAt),
			CreatedBy:    posUserEntity.CreatedBy.String(),
			UpdatedAt:    timestamppb.New(posUserEntity.UpdatedAt),
			UpdatedBy:    posUserEntity.UpdatedBy.String(),
		}

		// Store the user in Redis for future queries
		userData, err := json.Marshal(posUserEntity)
		if err != nil {
			return nil, err
		}
		err = r.redis.Set(context.Background(), userID, userData, 7*24*time.Hour).Err()
		if err != nil {
			return nil, err
		}

		return protoUser, nil
	} else if err != nil {
		return nil, err
	}

	// User found in Redis, unmarshal the data
	var posUserEntity entity.PosUser
	err = json.Unmarshal([]byte(userData), &posUserEntity)
	if err != nil {
		return nil, err
	}

	// Check if the UUID fields are nil before trying to convert them
	companyId := ""
	if posUserEntity.CompanyID != nil {
		companyId = posUserEntity.CompanyID.String()
	}

	branchId := ""
	if posUserEntity.BranchID != nil {
		branchId = posUserEntity.BranchID.String()
	}

	storeId := ""
	if posUserEntity.StoreID != nil {
		storeId = posUserEntity.StoreID.String()
	}

	// Convert entity.PosUser to pb.PosUser
	protoUser := &pb.PosUser{
		UserId:       posUserEntity.UserID.String(),
		Username:     posUserEntity.Username,
		PasswordHash: posUserEntity.PasswordHash,
		RoleId:       posUserEntity.RoleID.String(),
		CompanyId:    companyId,
		BranchId:     branchId,
		StoreId:      storeId,
		FirstName:    posUserEntity.FirstName,
		LastName:     posUserEntity.LastName,
		Email:        posUserEntity.Email,
		PhoneNumber:  posUserEntity.PhoneNumber,
		CreatedAt:    timestamppb.New(posUserEntity.CreatedAt),
		CreatedBy:    posUserEntity.CreatedBy.String(),
		UpdatedAt:    timestamppb.New(posUserEntity.UpdatedAt),
		UpdatedBy:    posUserEntity.UpdatedBy.String(),
	}

	return protoUser, nil
}

func (r *posUserRepository) ReadPosUserByUsername(username string) (*pb.PosUser, error) {
	// Scan the result into a PosUser entity
	var userEntity entity.PosUser
	if err := r.db.Where("username = ?", username).First(&userEntity).Error; err != nil {
		return nil, err
	}

	// Check if the UUID fields are nil before trying to convert them
	companyId := ""
	if userEntity.CompanyID != nil {
		companyId = userEntity.CompanyID.String()
	}

	branchId := ""
	if userEntity.BranchID != nil {
		branchId = userEntity.BranchID.String()
	}

	storeId := ""
	if userEntity.StoreID != nil {
		storeId = userEntity.StoreID.String()
	}

	// Convert the PosUser entity to a PosUser protobuf message
	user := &pb.PosUser{
		UserId:       userEntity.UserID.String(),
		Username:     userEntity.Username,
		PasswordHash: userEntity.PasswordHash,
		RoleId:       userEntity.RoleID.String(),
		CompanyId:    companyId,
		BranchId:     branchId,
		StoreId:      storeId,
		FirstName:    userEntity.FirstName,
		LastName:     userEntity.LastName,
		Email:        userEntity.Email,
		PhoneNumber:  userEntity.PhoneNumber,
		CreatedAt:    timestamppb.New(userEntity.CreatedAt),
		CreatedBy:    userEntity.CreatedBy.String(),
		UpdatedAt:    timestamppb.New(userEntity.UpdatedAt),
		UpdatedBy:    userEntity.UpdatedBy.String(),
	}

	return user, nil
}

func (r *posUserRepository) IsUsernameExist(username string) (bool, error) {
	var userEntity entity.PosUser
	if err := r.db.Where("username = ?", username).First(&userEntity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *posUserRepository) IsEmailExist(email string) (bool, error) {
	var userEntity entity.PosUser
	if err := r.db.Where("email = ?", email).First(&userEntity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *posUserRepository) UpdatePosUser(posUser *entity.PosUser) (*pb.PosUser, error) {
	if err := r.db.Save(posUser).Error; err != nil {
		return nil, err
	}

	// Check if the UUID fields are nil before trying to convert them
	companyId := ""
	if posUser.CompanyID != nil {
		companyId = posUser.CompanyID.String()
	}

	branchId := ""
	if posUser.BranchID != nil {
		branchId = posUser.BranchID.String()
	}

	storeId := ""
	if posUser.StoreID != nil {
		storeId = posUser.StoreID.String()
	}

	// Convert entity.PosUser to pb.PosUser
	protoUser := &pb.PosUser{
		UserId:       posUser.UserID.String(),
		Username:     posUser.Username,
		PasswordHash: posUser.PasswordHash,
		RoleId:       posUser.RoleID.String(),
		CompanyId:    companyId,
		BranchId:     branchId,
		StoreId:      storeId,
		FirstName:    posUser.FirstName,
		LastName:     posUser.LastName,
		Email:        posUser.Email,
		PhoneNumber:  posUser.PhoneNumber,
		CreatedAt:    timestamppb.New(posUser.CreatedAt),
		CreatedBy:    posUser.CreatedBy.String(),
		UpdatedAt:    timestamppb.New(posUser.UpdatedAt),
		UpdatedBy:    posUser.UpdatedBy.String(),
	}

	// Update the user in Redis
	userData, err := json.Marshal(posUser)
	if err != nil {
		return nil, err
	}
	err = r.redis.Set(context.Background(), protoUser.UserId, userData, 7*24*time.Hour).Err()
	if err != nil {
		return nil, err
	}

	return protoUser, nil
}

func (r *posUserRepository) DeletePosUser(userID string) error {
	if err := r.db.Where("user_id = ?", userID).Delete(&entity.PosUser{}).Error; err != nil {
		return err
	}

	// Delete the user from Redis
	err := r.redis.Del(context.Background(), userID).Err()
	if err != nil {
		return err
	}

	return nil
}
