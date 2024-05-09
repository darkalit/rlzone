package items

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, item *Item) error
	Get(ctx context.Context, query *GetItemQuery) ([]Item, error)
	GetById(ctx context.Context, id uint) (*Item, error)
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

	err := tx.Create(&item).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemsRepo) Get(ctx context.Context, query *GetItemQuery) ([]Item, error) {
	var items []Item
	tx := r.db.WithContext(ctx)
	tx = tx.Model(&Item{}).Preload("Stock")

	if id := query.ID; id != 0 {
		tx = tx.Where("id = ?", id)
	}

	if quality := query.Quality; quality != "" {
		quality = "%" + quality + "%"
		tx = tx.Where("quality = ?", quality)
	}

	// if maxprice := query.MaxPrice; maxprice != 0 {
	// 	tx = tx.Where("price <= ?", maxprice)
	// }
	//
	// if minprice := query.MinPrice; minprice != 0 {
	// 	tx = tx.Where("price >= ?", minprice)
	// }

	pagesize := query.PageSize
	if pagesize != 0 {
		tx = tx.Limit(pagesize)
	}

	if page := query.Page; page != 0 {
		tx = tx.Offset((page - 1) * pagesize)
	}

	err := tx.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
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
