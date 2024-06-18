package repository

import (
	"errors"
	"math"
	"os"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/entity"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PosStoreBranchRepository interface {
	CreatePosStoreBranch(posStoreBranch *entity.PosStoreBranch) error
	ReadPosStoreBranch(branchID string) (*pb.PosStoreBranch, error)
	UpdatePosStoreBranch(posStoreBranch *entity.PosStoreBranch) (*pb.PosStoreBranch, error)
	DeletePosStoreBranch(branchID string) error
	ReadAllPosStoreBranches(pagination dto.Pagination, roleName string, jwtPayload *pb.JWTPayload) (*dto.PaginationResult, error)
}

type posStoreBranchRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosStoreBranchRepository(db *gorm.DB, redis *redis.Client) PosStoreBranchRepository {
	return &posStoreBranchRepository{
		db:    db,
		redis: redis,
	}
}

func (r *posStoreBranchRepository) CreatePosStoreBranch(posStoreBranch *entity.PosStoreBranch) error {
	result := r.db.Create(posStoreBranch)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *posStoreBranchRepository) ReadAllPosStoreBranches(pagination dto.Pagination, roleName string, jwtPayload *pb.JWTPayload) (*dto.PaginationResult, error) {
	var posStoreBranches []entity.PosStoreBranch
	var totalRecords int64

	query := r.db.Model(&entity.PosStoreBranch{})

	companyRole := os.Getenv("COMPANY_USER_ROLE")

	switch roleName {

	case companyRole:
		query = query.Where("company_id = ?", jwtPayload.CompanyId)
	default:
		return nil, errors.New("invalid role")
	}

	if pagination.Limit > 0 && pagination.Page > 0 {
		offset := (pagination.Page - 1) * pagination.Limit
		query = query.Offset(offset).Limit(pagination.Limit)
	}

	query.Find(&posStoreBranches)
	query.Count(&totalRecords)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pagination.Limit)))

	return &dto.PaginationResult{
		TotalRecords: totalRecords,
		Records:      posStoreBranches,
		CurrentPage:  pagination.Page,
		TotalPages:   totalPages,
	}, nil
}

func (r *posStoreBranchRepository) ReadPosStoreBranch(branchID string) (*pb.PosStoreBranch, error) {
	var posStoreBranchEntity entity.PosStoreBranch

	if err := r.db.Where("branch_id = ?", branchID).First(&posStoreBranchEntity).Error; err != nil {
		return nil, err
	}

	// Convert entity.PosStoreBranch to pb.PosStoreBranch
	posStoreBranch := &pb.PosStoreBranch{
		BranchId:   posStoreBranchEntity.BranchID.String(),
		BranchName: posStoreBranchEntity.BranchName,
		CompanyId:  posStoreBranchEntity.CompanyID.String(),
		CreatedAt:  timestamppb.New(posStoreBranchEntity.CreatedAt),
		CreatedBy:  posStoreBranchEntity.CreatedBy.String(),
		UpdatedAt:  timestamppb.New(posStoreBranchEntity.UpdatedAt),
		UpdatedBy:  posStoreBranchEntity.UpdatedBy.String(),
	}

	return posStoreBranch, nil
}

func (r *posStoreBranchRepository) UpdatePosStoreBranch(posStoreBranch *entity.PosStoreBranch) (*pb.PosStoreBranch, error) {

	if err := r.db.Save(posStoreBranch).Error; err != nil {
		return nil, err
	}

	// Convert updated entity.PosStoreBranch back to pb.PosStoreBranch
	updatedPosStoreBranch := &pb.PosStoreBranch{
		BranchId:   posStoreBranch.BranchID.String(),
		BranchName: posStoreBranch.BranchName,
		CompanyId:  posStoreBranch.CompanyID.String(),
		CreatedAt:  timestamppb.New(posStoreBranch.CreatedAt),
		CreatedBy:  posStoreBranch.CreatedBy.String(),
		UpdatedAt:  timestamppb.New(posStoreBranch.UpdatedAt),
		UpdatedBy:  posStoreBranch.UpdatedBy.String(),
	}

	return updatedPosStoreBranch, nil
}

func (r *posStoreBranchRepository) DeletePosStoreBranch(branchID string) error {
	if err := r.db.Where("branch_id = ?", branchID).Delete(&entity.PosStoreBranch{}).Error; err != nil {
		return err
	}
	return nil
}
