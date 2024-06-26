package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

type PosStoreRepository interface {
	CreatePosStore(posStore *entity.PosStore) error
	ReadPosStore(storeID string) (*pb.PosStore, error)
	UpdatePosStore(posStore *entity.PosStore) (*pb.PosStore, error)
	DeletePosStore(storeID string) error
	ReadAllPosStores(pagination dto.Pagination, roleName string, jwtPayload *pb.JWTPayload) (*dto.PaginationResult, error)
	GetNextReceiptID(storeID string) (int, error)
}

type posStoreRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosStoreRepository(db *gorm.DB, redis *redis.Client) PosStoreRepository {
	return &posStoreRepository{
		db:    db,
		redis: redis,
	}
}

func (r *posStoreRepository) CreatePosStore(posStore *entity.PosStore) error {
	result := r.db.Create(posStore)
	if result.Error != nil {
		return result.Error
	}

	// Create a new sequence for the store
	seqName := fmt.Sprintf("receipt_id_seq_store_s%s", posStore.StoreID)
	var count int64
	row := r.db.Raw("SELECT COUNT(*) FROM pg_sequences WHERE sequencename = ?", seqName).Row()
	row.Scan(&count)
	if count == 0 {
		db := r.db.Exec(fmt.Sprintf("CREATE SEQUENCE \"%s\" START 1000", seqName))
		if db.Error != nil {
			return db.Error
		}
	}

	return nil
}

func (r *posStoreRepository) ReadAllPosStores(pagination dto.Pagination, roleName string, jwtPayload *pb.JWTPayload) (*dto.PaginationResult, error) {
	var posStores []entity.PosStore
	var totalRecords int64

	query := r.db.Model(&entity.PosStore{})

	companyRole := os.Getenv("COMPANY_USER_ROLE")
	branchRole := os.Getenv("BRANCH_USER_ROLE")

	switch roleName {
	case companyRole:
		query = query.Where("company_id = ?", jwtPayload.CompanyId)
	case branchRole:
		query = query.Where("branch_id = ?", jwtPayload.BranchId)
	default:
		return nil, errors.New("invalid role")
	}

	if pagination.Limit > 0 && pagination.Page > 0 {
		offset := (pagination.Page - 1) * pagination.Limit
		query = query.Offset(offset).Limit(pagination.Limit)
	}

	query.Find(&posStores)
	query.Count(&totalRecords)

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pagination.Limit)))

	return &dto.PaginationResult{
		TotalRecords: totalRecords,
		Records:      posStores,
		CurrentPage:  pagination.Page,
		TotalPages:   totalPages,
	}, nil
}

func (r *posStoreRepository) ReadPosStore(storeID string) (*pb.PosStore, error) {
	// Try to get the store from Redis first
	storeData, err := r.redis.Get(context.Background(), storeID).Result()
	if err == redis.Nil {
		// Store not found in Redis, get from PostgreSQL
		var posStoreEntity entity.PosStore
		if err := r.db.Where("store_id = ?", storeID).First(&posStoreEntity).Error; err != nil {
			return nil, err
		}

		// Convert entity.PosStore to pb.PosStore
		posStore := &pb.PosStore{
			StoreId:   posStoreEntity.StoreID.String(),
			StoreName: posStoreEntity.StoreName,
			BranchId:  posStoreEntity.BranchID.String(),
			Location:  posStoreEntity.Location,
			CompanyId: posStoreEntity.CompanyID.String(),
			CreatedAt: timestamppb.New(posStoreEntity.CreatedAt),
			CreatedBy: posStoreEntity.CreatedBy.String(),
			UpdatedAt: timestamppb.New(posStoreEntity.UpdatedAt),
			UpdatedBy: posStoreEntity.UpdatedBy.String(),
		}

		// Store the store in Redis for future queries
		storeData, err := json.Marshal(posStoreEntity)
		if err != nil {
			return nil, err
		}
		err = r.redis.Set(context.Background(), storeID, storeData, 7*24*time.Hour).Err()
		if err != nil {
			return nil, err
		}

		return posStore, nil
	} else if err != nil {
		return nil, err
	}

	// Store found in Redis, unmarshal the data
	var posStoreEntity entity.PosStore
	err = json.Unmarshal([]byte(storeData), &posStoreEntity)
	if err != nil {
		return nil, err
	}

	// Convert entity.PosStore to pb.PosStore
	posStore := &pb.PosStore{
		StoreId:   posStoreEntity.StoreID.String(),
		StoreName: posStoreEntity.StoreName,
		BranchId:  posStoreEntity.BranchID.String(),
		Location:  posStoreEntity.Location,
		CompanyId: posStoreEntity.CompanyID.String(),
		CreatedAt: timestamppb.New(posStoreEntity.CreatedAt),
		CreatedBy: posStoreEntity.CreatedBy.String(),
		UpdatedAt: timestamppb.New(posStoreEntity.UpdatedAt),
		UpdatedBy: posStoreEntity.UpdatedBy.String(),
	}

	return posStore, nil
}

func (r *posStoreRepository) UpdatePosStore(posStore *entity.PosStore) (*pb.PosStore, error) {
	if err := r.db.Save(posStore).Error; err != nil {
		return nil, err
	}

	// Convert updated entity.PosStore back to pb.PosStore
	updatedPosStore := &pb.PosStore{
		StoreId:   posStore.StoreID.String(),
		StoreName: posStore.StoreName,
		BranchId:  posStore.BranchID.String(),
		Location:  posStore.Location,
		CompanyId: posStore.CompanyID.String(),
		CreatedAt: timestamppb.New(posStore.CreatedAt),
		CreatedBy: posStore.CreatedBy.String(),
		UpdatedAt: timestamppb.New(posStore.UpdatedAt),
		UpdatedBy: posStore.UpdatedBy.String(),
	}

	// Update the store in Redis
	storeData, err := json.Marshal(posStore)
	if err != nil {
		return nil, err
	}
	err = r.redis.Set(context.Background(), updatedPosStore.StoreId, storeData, 7*24*time.Hour).Err()
	if err != nil {
		return nil, err
	}

	return updatedPosStore, nil
}

func (r *posStoreRepository) DeletePosStore(storeID string) error {
	if err := r.db.Where("store_id = ?", storeID).Delete(&entity.PosStore{}).Error; err != nil {
		return err
	}

	// Delete the store from Redis
	err := r.redis.Del(context.Background(), storeID).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *posStoreRepository) GetNextReceiptID(storeID string) (int, error) {
	seqName := fmt.Sprintf("receipt_id_seq_store_s%s", storeID)
	var receiptID int
	row := r.db.Raw(fmt.Sprintf("SELECT nextval('\"%s\"')", seqName)).Row()
	err := row.Scan(&receiptID)
	if err != nil {
		return 0, err
	}
	return receiptID, nil
}
