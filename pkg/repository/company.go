// package repository

// import (
// 	"errors"
// 	"math"

// 	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
// 	"github.com/Andrewalifb/alpha-pos-system/company-service/dto"
// 	"github.com/Andrewalifb/alpha-pos-system/company-service/entity"
// 	"github.com/go-redis/redis/v8"
// 	"github.com/jinzhu/gorm"
// 	"google.golang.org/protobuf/types/known/timestamppb"
// )

// type PosCompanyRepository interface {
// 	CreatePosCompany(posCompany *entity.PosCompany) error
// 	ReadPosCompany(companyID string) (*pb.PosCompany, error)
// 	UpdatePosCompany(posCompany *entity.PosCompany) (*pb.PosCompany, error)
// 	DeletePosCompany(companyID string) error
// 	ReadAllPosCompanies(pagination dto.Pagination, roleName string, jwtPayload *pb.JWTPayload) (*dto.PaginationResult, error)
// }

// type posCompanyRepository struct {
// 	db    *gorm.DB
// 	redis *redis.Client
// }

// func NewPosCompanyRepository(db *gorm.DB, redis *redis.Client) PosCompanyRepository {
// 	return &posCompanyRepository{
// 		db:    db,
// 		redis: redis,
// 	}
// }
// func (r *posCompanyRepository) CreatePosCompany(posCompany *entity.PosCompany) error {
// 	result := r.db.Create(posCompany)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func (r *posCompanyRepository) ReadAllPosCompanies(pagination dto.Pagination, roleName string, jwtPayload *pb.JWTPayload) (*dto.PaginationResult, error) {
// 	var posCompanies []entity.PosCompany
// 	var totalRecords int64

// 	query := r.db.Model(&entity.PosCompany{})

// 	switch roleName {
// 	case "super user":
// 		// No filters
// 	case "company":
// 		query = query.Where("company_id = ?", jwtPayload.CompanyId)
// 	case "branch":
// 		query = query.Where("branch_id = ?", jwtPayload.BranchId)
// 	case "store":
// 		return nil, errors.New("store users are not allowed to retrieve roles")
// 	default:
// 		return nil, errors.New("invalid role")
// 	}

// 	if pagination.Limit > 0 && pagination.Page > 0 {
// 		offset := (pagination.Page - 1) * pagination.Limit
// 		query = query.Offset(offset).Limit(pagination.Limit)
// 	}

// 	query.Find(&posCompanies)
// 	query.Count(&totalRecords)

// 	totalPages := int(math.Ceil(float64(totalRecords) / float64(pagination.Limit)))

// 	return &dto.PaginationResult{
// 		TotalRecords: totalRecords,
// 		Records:      posCompanies,
// 		CurrentPage:  pagination.Page,
// 		TotalPages:   totalPages,
// 	}, nil
// }

// func (r *posCompanyRepository) ReadPosCompany(companyID string) (*pb.PosCompany, error) {
// 	var posCompany entity.PosCompany
// 	if err := r.db.Where("company_id = ?", companyID).First(&posCompany).Error; err != nil {
// 		return nil, err
// 	}

// 	protoCompany := &pb.PosCompany{
// 		CompanyId:   posCompany.CompanyID.String(),
// 		CompanyName: posCompany.CompanyName,
// 		CreatedAt:   timestamppb.New(posCompany.CreatedAt),
// 		CreatedBy:   posCompany.CreatedBy.String(),
// 		UpdatedAt:   timestamppb.New(posCompany.UpdatedAt),
// 		UpdatedBy:   posCompany.UpdatedBy.String(),
// 	}

// 	return protoCompany, nil
// }

// func (r *posCompanyRepository) UpdatePosCompany(posCompany *entity.PosCompany) (*pb.PosCompany, error) {
// 	if err := r.db.Save(posCompany).Error; err != nil {
// 		return nil, err
// 	}
// 	protoCompany := &pb.PosCompany{
// 		CompanyId:   posCompany.CompanyID.String(),
// 		CompanyName: posCompany.CompanyName,
// 		CreatedAt:   timestamppb.New(posCompany.CreatedAt),
// 		CreatedBy:   posCompany.CreatedBy.String(),
// 		UpdatedAt:   timestamppb.New(posCompany.UpdatedAt),
// 		UpdatedBy:   posCompany.UpdatedBy.String(),
// 	}

// 	return protoCompany, nil
// }

// func (r *posCompanyRepository) DeletePosCompany(companyID string) error {
// 	if err := r.db.Where("company_id = ?", companyID).Delete(&pb.PosCompany{}).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

package repository

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"time"

	pb "github.com/Andrewalifb/alpha-pos-system-company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/dto"
	"github.com/Andrewalifb/alpha-pos-system-company-service/entity"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PosCompanyRepository interface {
	CreatePosCompany(posCompany *entity.PosCompany) error
	ReadPosCompany(companyID string) (*pb.PosCompany, error)
	UpdatePosCompany(posCompany *entity.PosCompany) (*pb.PosCompany, error)
	DeletePosCompany(companyID string) error
	ReadAllPosCompanies(pagination dto.Pagination, roleName string, jwtPayload *pb.JWTPayload) (*dto.PaginationResult, error)
}

type posCompanyRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosCompanyRepository(db *gorm.DB, redis *redis.Client) PosCompanyRepository {
	return &posCompanyRepository{
		db:    db,
		redis: redis,
	}
}

func (r *posCompanyRepository) CreatePosCompany(posCompany *entity.PosCompany) error {
	result := r.db.Create(posCompany)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *posCompanyRepository) ReadAllPosCompanies(pagination dto.Pagination, roleName string, jwtPayload *pb.JWTPayload) (*dto.PaginationResult, error) {
	var posCompanies []entity.PosCompany
	var totalRecords int64

	query := r.db.Model(&entity.PosCompany{})

	switch roleName {
	case "super user":
		// No filters
	case "company":
		query = query.Where("company_id = ?", jwtPayload.CompanyId)
	case "branch":
		return nil, errors.New("branch users are not allowed to retrieve roles")
	case "store":
		return nil, errors.New("store users are not allowed to retrieve roles")
	default:
		return nil, errors.New("invalid role")
	}

	if pagination.Limit > 0 && pagination.Page > 0 {
		offset := (pagination.Page - 1) * pagination.Limit
		query = query.Offset(offset).Limit(pagination.Limit)
	}

	query.Find(&posCompanies)
	query.Count(&totalRecords)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pagination.Limit)))

	return &dto.PaginationResult{
		TotalRecords: totalRecords,
		Records:      posCompanies,
		CurrentPage:  pagination.Page,
		TotalPages:   totalPages,
	}, nil
}

func (r *posCompanyRepository) ReadPosCompany(companyID string) (*pb.PosCompany, error) {
	// Try to get the company from Redis first
	companyData, err := r.redis.Get(context.Background(), companyID).Result()
	if err == redis.Nil {
		// Company not found in Redis, get from PostgreSQL
		var posCompanyEntity entity.PosCompany
		if err := r.db.Where("company_id = ?", companyID).First(&posCompanyEntity).Error; err != nil {
			return nil, err
		}

		// Convert entity.PosCompany to pb.PosCompany
		posCompany := &pb.PosCompany{
			CompanyId:   posCompanyEntity.CompanyID.String(),
			CompanyName: posCompanyEntity.CompanyName,
			CreatedAt:   timestamppb.New(posCompanyEntity.CreatedAt),
			CreatedBy:   posCompanyEntity.CreatedBy.String(),
			UpdatedAt:   timestamppb.New(posCompanyEntity.UpdatedAt),
			UpdatedBy:   posCompanyEntity.UpdatedBy.String(),
		}

		// Store the company in Redis for future queries
		companyData, err := json.Marshal(posCompanyEntity)
		if err != nil {
			return nil, err
		}
		err = r.redis.Set(context.Background(), companyID, companyData, 7*24*time.Hour).Err()
		if err != nil {
			return nil, err
		}

		return posCompany, nil
	} else if err != nil {
		return nil, err
	}

	// Company found in Redis, unmarshal the data
	var posCompanyEntity entity.PosCompany
	err = json.Unmarshal([]byte(companyData), &posCompanyEntity)
	if err != nil {
		return nil, err
	}

	// Convert entity.PosCompany to pb.PosCompany
	posCompany := &pb.PosCompany{
		CompanyId:   posCompanyEntity.CompanyID.String(),
		CompanyName: posCompanyEntity.CompanyName,
		CreatedAt:   timestamppb.New(posCompanyEntity.CreatedAt),
		CreatedBy:   posCompanyEntity.CreatedBy.String(),
		UpdatedAt:   timestamppb.New(posCompanyEntity.UpdatedAt),
		UpdatedBy:   posCompanyEntity.UpdatedBy.String(),
	}

	return posCompany, nil
}

func (r *posCompanyRepository) UpdatePosCompany(posCompany *entity.PosCompany) (*pb.PosCompany, error) {
	if err := r.db.Save(posCompany).Error; err != nil {
		return nil, err
	}

	// Convert updated entity.PosCompany back to pb.PosCompany
	updatedPosCompany := &pb.PosCompany{
		CompanyId:   posCompany.CompanyID.String(),
		CompanyName: posCompany.CompanyName,
		CreatedAt:   timestamppb.New(posCompany.CreatedAt),
		CreatedBy:   posCompany.CreatedBy.String(),
		UpdatedAt:   timestamppb.New(posCompany.UpdatedAt),
		UpdatedBy:   posCompany.UpdatedBy.String(),
	}

	// Update the company in Redis
	companyData, err := json.Marshal(posCompany)
	if err != nil {
		return nil, err
	}
	err = r.redis.Set(context.Background(), updatedPosCompany.CompanyId, companyData, 7*24*time.Hour).Err()
	if err != nil {
		return nil, err
	}

	return updatedPosCompany, nil
}

func (r *posCompanyRepository) DeletePosCompany(companyID string) error {
	if err := r.db.Where("company_id = ?", companyID).Delete(&pb.PosCompany{}).Error; err != nil {
		return err
	}

	// Delete the company from Redis
	err := r.redis.Del(context.Background(), companyID).Err()
	if err != nil {
		return err
	}

	return nil
}
