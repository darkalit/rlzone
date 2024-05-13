package items

import (
	"context"

	"gorm.io/gorm"

	"github.com/darkalit/rlzone/server/pkg/pagination"
)

type Repository interface {
	Create(ctx context.Context, item *Item) error
	Get(ctx context.Context, query *GetItemsQuery) (*GetResponse, error)
	GetById(ctx context.Context, id uint) (*Item, error)

	CreateStock(ctx context.Context, stock *Stock) error
}

type ItemsRepo struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemsRepo {
	return &ItemsRepo{
		db: db,
	}
}

func (r *ItemsRepo) Create(ctx context.Context, item *Item) error {
	tx := r.db.WithContext(ctx)

	err := tx.Create(item).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemsRepo) Get(ctx context.Context, query *GetItemsQuery) (*GetResponse, error) {
	var totalCount int64
	var items []Item
	tx := r.db.WithContext(ctx)
	tx = tx.Model(&Item{}).Preload("Stock")

	if id := query.ID; id != 0 {
		tx = tx.Where("id = ?", id)
	}

	if itemType := query.Type; itemType != "" {
		itemType = "%" + itemType + "%"
		tx = tx.Where("type = ?", itemType)
	}

	if quality := query.Quality; quality != "" {
		quality = "%" + quality + "%"
		tx = tx.Where("quality = ?", quality)
	}

	// if onlyInStock := query.OnlyInStock; onlyInStock {
	// 	tx = tx.Where("stock")
	// }

	// if maxprice := query.MaxPrice; maxprice != 0 {
	// 	tx = tx.Where("price <= ?", maxprice)
	// }
	//
	// if minprice := query.MinPrice; minprice != 0 {
	// 	tx = tx.Where("price >= ?", minprice)
	// }

	tx.Count(&totalCount)

	pagesize := query.PageSize
	page := query.Page

	pagination.SanitizePages(&page, &pagesize)
	offset, limit := pagination.GetOffsetAndLimit(page, pagesize)

	tx = tx.Limit(limit)
	tx = tx.Offset(offset)

	err := tx.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return &GetResponse{
		Items:      items,
		Pagination: pagination.GetPagination(int64(page), int64(pagesize), totalCount),
	}, nil
}

func (r *ItemsRepo) GetById(ctx context.Context, id uint) (*Item, error) {
	var item Item

	tx := r.db.WithContext(ctx)
	err := tx.Preload("Stock").First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemsRepo) CreateStock(ctx context.Context, stock *Stock) error {
	tx := r.db.WithContext(ctx)

	err := tx.Create(stock).Error
	if err != nil {
		return err
	}

	return nil
}
