package repository

import (
	"context"
	"encoding/json"
	"math"
	"time"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/entity"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PosRoleRepository interface {
	CreatePosRole(posRole *entity.PosRole) error
	ReadPosRole(roleID string) (*pb.PosRole, error)
	UpdatePosRole(posRole *entity.PosRole) (*pb.PosRole, error)
	DeletePosRole(roleID string) error
	ReadAllPosRoles(pagination dto.Pagination, roleName string, jwtPayload *pb.JWTPayload) (*dto.PaginationResult, error)
}

type posRoleRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosRoleRepository(db *gorm.DB, redis *redis.Client) PosRoleRepository {
	return &posRoleRepository{
		db:    db,
		redis: redis,
	}
}

func (r *posRoleRepository) CreatePosRole(posRole *entity.PosRole) error {
	result := r.db.Create(posRole)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *posRoleRepository) ReadAllPosRoles(pagination dto.Pagination, roleName string, jwtPayload *pb.JWTPayload) (*dto.PaginationResult, error) {
	var posRoles []entity.PosRole
	var totalRecords int64

	query := r.db.Model(&entity.PosRole{})

	if pagination.Limit > 0 && pagination.Page > 0 {
		offset := (pagination.Page - 1) * pagination.Limit
		query = query.Offset(offset).Limit(pagination.Limit)
	}

	query.Find(&posRoles)
	query.Count(&totalRecords)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pagination.Limit)))

	return &dto.PaginationResult{
		TotalRecords: totalRecords,
		Records:      posRoles,
		CurrentPage:  pagination.Page,
		TotalPages:   totalPages,
	}, nil
}

func (r *posRoleRepository) ReadPosRole(roleID string) (*pb.PosRole, error) {
	// Try to get the role from Redis first
	roleData, err := r.redis.Get(context.Background(), roleID).Result()

	if err == redis.Nil {

		// Role not found in Redis, get from PostgreSQL
		var posRoleEntity entity.PosRole
		if err := r.db.Where("role_id = ?", roleID).First(&posRoleEntity).Error; err != nil {
			return nil, err
		}

		// Convert entity.PosRole to pb.PosRole
		posRole := &pb.PosRole{
			RoleId:    posRoleEntity.RoleID.String(),
			RoleName:  posRoleEntity.RoleName,
			CreatedAt: timestamppb.New(posRoleEntity.CreatedAt),
			CreatedBy: posRoleEntity.CreatedBy.String(),
			UpdatedAt: timestamppb.New(posRoleEntity.UpdatedAt),
			UpdatedBy: posRoleEntity.UpdatedBy.String(),
		}

		// Store the role in Redis for future queries
		roleData, err := json.Marshal(posRoleEntity)
		if err != nil {
			return nil, err
		}
		err = r.redis.Set(context.Background(), roleID, roleData, 7*24*time.Hour).Err()
		if err != nil {
			return nil, err
		}

		return posRole, nil
	} else if err != nil {
		return nil, err
	}

	// Role found in Redis, unmarshal the data
	var posRoleEntity entity.PosRole
	err = json.Unmarshal([]byte(roleData), &posRoleEntity)
	if err != nil {
		return nil, err
	}

	// Convert entity.PosRole to pb.PosRole
	posRole := &pb.PosRole{
		RoleId:    posRoleEntity.RoleID.String(),
		RoleName:  posRoleEntity.RoleName,
		CreatedAt: timestamppb.New(posRoleEntity.CreatedAt),
		CreatedBy: posRoleEntity.CreatedBy.String(),
		UpdatedAt: timestamppb.New(posRoleEntity.UpdatedAt),
		UpdatedBy: posRoleEntity.UpdatedBy.String(),
	}

	return posRole, nil
}

func (r *posRoleRepository) UpdatePosRole(posRole *entity.PosRole) (*pb.PosRole, error) {
	if err := r.db.Save(posRole).Error; err != nil {
		return nil, err
	}

	// Convert updated entity.PosRole back to pb.PosRole
	updatedPosRole := &pb.PosRole{
		RoleId:    posRole.RoleID.String(),
		RoleName:  posRole.RoleName,
		CreatedAt: timestamppb.New(posRole.CreatedAt),
		CreatedBy: posRole.CreatedBy.String(),
		UpdatedAt: timestamppb.New(posRole.UpdatedAt),
		UpdatedBy: posRole.UpdatedBy.String(),
	}

	// Update the role in Redis
	roleData, err := json.Marshal(posRole)
	if err != nil {
		return nil, err
	}
	err = r.redis.Set(context.Background(), updatedPosRole.RoleId, roleData, 7*24*time.Hour).Err()
	if err != nil {
		return nil, err
	}

	return updatedPosRole, nil
}

func (r *posRoleRepository) DeletePosRole(roleID string) error {
	if err := r.db.Where("role_id = ?", roleID).Delete(&entity.PosRole{}).Error; err != nil {
		return err
	}

	// Delete the role from Redis
	err := r.redis.Del(context.Background(), roleID).Err()
	if err != nil {
		return err
	}

	return nil
}
