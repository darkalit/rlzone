package items

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/internal/users"
	"github.com/darkalit/rlzone/server/pkg/httpErrors"
	"github.com/darkalit/rlzone/server/pkg/pagination"
)

type Repository interface {
	Create(ctx context.Context, item *Item) error
	Get(ctx context.Context, query *GetItemsQuery) (*GetResponse, error)
	GetById(ctx context.Context, id uint) (*Item, error)

	CreateStock(ctx context.Context, stock *Stock) error
	GetInventoryItem(ctx context.Context, inventoryItem *InventoryItem) error
	BuyItem(ctx context.Context, itemId uint, userId uint) (*InventoryItem, error)
	SellItem(ctx context.Context, itemId uint, userId uint) (*InventoryItem, error)
	GetInventory(
		ctx context.Context,
		query *GetItemsQuery,
		userId uint,
	) (*GetInventoryResponse, error)
}

type ItemsRepo struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewItemRepository(cfg *config.Config, db *gorm.DB) *ItemsRepo {
	return &ItemsRepo{
		cfg: cfg,
		db:  db,
	}
}

func (r *ItemsRepo) Create(ctx context.Context, item *Item) error {
	tx := r.db.WithContext(ctx)
	tx = tx.Exec(fmt.Sprintf("USE %s", r.cfg.DBName))

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
	tx = tx.Exec(fmt.Sprintf("USE %s", r.cfg.DBName))
	tx = tx.Model(&Item{}).Preload("Stock")

	if id := query.ID; id != 0 {
		tx = tx.Where("id = ?", id)
	}

	if itemType := query.Type; itemType != "" {
		tx = tx.Where("type = ?", itemType)
	}

	if quality := query.Quality; quality != "" {
		tx = tx.Where("quality = ?", quality)
	}

	if query.OnlyInStock || query.MaxPrice != 0 || query.MinPrice != 0 {
		tx = tx.InnerJoins("Stock")
	}

	if maxprice := query.MaxPrice; maxprice != 0 {
		tx = tx.Where("`Stock`.`price` <= ?", maxprice)
	}

	if minprice := query.MinPrice; minprice != 0 {
		tx = tx.Where("`Stock`.`price` >= ?", minprice)
	}

	if name := query.Name; name != "" {
		name = "%" + name + "%"
		tx = tx.Where("name LIKE ?", name)
	}

	if sort := query.Sort; sort != "" {
		var sortString string
		switch sort {
		case "id_desc":
			sortString = "id desc"
		case "id_asc":
			sortString = "id asc"
		case "price_id_desc":
			sortString = "`Stock`.`price` desc"
		case "price_id_asc":
			sortString = "`Stock`.`price` asc"
		case "quality_desc":
			sortString = `CASE
        WHEN quality = 'Common' THEN 0
        WHEN quality = 'Uncommon' THEN 1
        WHEN quality = 'Limited' THEN 2
        WHEN quality = 'Rare' THEN 3
        WHEN quality = 'Very Rare' THEN 4
        WHEN quality = 'Import' THEN 5
        WHEN quality = 'Exotic' THEN 6
        WHEN quality = 'Black Market' THEN 7
        WHEN quality = 'Premium' THEN 8
        WHEN quality = 'Legacy' THEN 9
        ELSE 99
      END desc`
		case "quality_asc":
			sortString = `CASE
        WHEN quality = 'Common' THEN 0
        WHEN quality = 'Uncommon' THEN 1
        WHEN quality = 'Limited' THEN 2
        WHEN quality = 'Rare' THEN 3
        WHEN quality = 'Very Rare' THEN 4
        WHEN quality = 'Import' THEN 5
        WHEN quality = 'Exotic' THEN 6
        WHEN quality = 'Black Market' THEN 7
        WHEN quality = 'Premium' THEN 8
        WHEN quality = 'Legacy' THEN 9
        ELSE 99
      END asc`
		default:
			sortString = ""
		}

		tx = tx.Order(sortString)
	}

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
	tx = tx.Exec(fmt.Sprintf("USE %s", r.cfg.DBName))
	err := tx.Preload("Stock").First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ItemsRepo) CreateStock(ctx context.Context, stock *Stock) error {
	tx := r.db.WithContext(ctx)
	tx = tx.Exec(fmt.Sprintf("USE %s", r.cfg.DBName))

	err := tx.Create(stock).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemsRepo) UpdateOrCreateInventoryItem(
	ctx context.Context,
	inventoryItem *InventoryItem,
) error {
	tx := r.db.WithContext(ctx)
	tx = tx.Exec(fmt.Sprintf("USE %s", r.cfg.DBName))

	err := tx.Save(inventoryItem).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemsRepo) GetInventoryItem(ctx context.Context, inventoryItem *InventoryItem) error {
	tx := r.db.WithContext(ctx)
	tx = tx.Exec(fmt.Sprintf("USE %s", r.cfg.DBName))

	err := tx.First(inventoryItem).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ItemsRepo) BuyItem(ctx context.Context, itemId uint, userId uint) (*InventoryItem, error) {
	var item Item
	var inventoryItem InventoryItem
	var user users.User
	tx := r.db.WithContext(ctx)
	tx = tx.Exec(fmt.Sprintf("USE %s", r.cfg.DBName)).Begin()

	item = Item{ID: itemId}
	err := tx.First(&item).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.Stock == nil {
		tx.Rollback()
		return nil, httpErrors.NewRestErrorMessage(400, "Not available to buy")
	}

	inventoryItem = InventoryItem{ItemID: itemId, UserID: userId}
	err = tx.FirstOrCreate(&inventoryItem).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	inventoryItem.Count++
	err = tx.Save(&inventoryItem).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	user = users.User{ID: userId}
	err = tx.First(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if user.Balance < item.Stock.Price {
		tx.Rollback()
		return nil, httpErrors.NewRestErrorMessage(400, "No money")
	}

	user.Balance -= item.Stock.Price

	err = tx.Save(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return &inventoryItem, nil
}

func (r *ItemsRepo) SellInventoryItem(
	ctx context.Context,
	itemId uint,
	userId uint,
) (*InventoryItem, error) {
	var item Item
	var inventoryItem InventoryItem
	var user users.User
	tx := r.db.WithContext(ctx)
	tx = tx.Exec(fmt.Sprintf("USE %s", r.cfg.DBName)).Begin()

	inventoryItem = InventoryItem{ItemID: itemId, UserID: userId}
	err := tx.FirstOrCreate(&inventoryItem).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if inventoryItem.Count == 0 {
		tx.Rollback()
		return nil, httpErrors.NewRestErrorMessage(400, "No item in inventory")
	}

	inventoryItem.Count--
	err = tx.Save(&inventoryItem).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	item = Item{ID: itemId}
	err = tx.First(&item).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if item.Stock == nil {
		tx.Rollback()
		return nil, httpErrors.NewRestErrorMessage(400, "Not available to sell")
	}

	user = users.User{ID: userId}
	err = tx.First(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	user.Balance += item.Stock.Price

	err = tx.Save(&user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return &inventoryItem, nil
}

func (r *ItemsRepo) GetInventory(
	ctx context.Context,
	query *GetItemsQuery,
	userId uint,
) (*GetInventoryResponse, error) {
	var totalCount int64
	var inventoryItems []InventoryItem
	tx := r.db.WithContext(ctx)
	tx = tx.Exec(fmt.Sprintf("USE %s", r.cfg.DBName))
	tx = tx.Model(&InventoryItem{}).Preload("Item")

	tx = tx.Where(&InventoryItem{UserID: userId})

	if id := query.ID; id != 0 {
		tx = tx.Where("id = ?", id)
	}

	if itemType := query.Type; itemType != "" {
		tx = tx.Where("type = ?", itemType)
	}

	if quality := query.Quality; quality != "" {
		tx = tx.Where("quality = ?", quality)
	}

	if query.OnlyInStock || query.MaxPrice != 0 || query.MinPrice != 0 {
		tx = tx.InnerJoins("Stock")
	}

	if maxprice := query.MaxPrice; maxprice != 0 {
		tx = tx.Where("`Stock`.`price` <= ?", maxprice)
	}

	if minprice := query.MinPrice; minprice != 0 {
		tx = tx.Where("`Stock`.`price` >= ?", minprice)
	}

	if name := query.Name; name != "" {
		name = "%" + name + "%"
		tx = tx.Where("name LIKE ?", name)
	}

	tx = tx.Where("count > 0")

	if sort := query.Sort; sort != "" {
		var sortString string
		switch sort {
		case "id_desc":
			sortString = "id desc"
		case "id_asc":
			sortString = "id asc"
		case "price_id_desc":
			sortString = "`Stock`.`price` desc"
		case "price_id_asc":
			sortString = "`Stock`.`price` asc"
		case "quality_desc":
			sortString = `CASE
        WHEN quality = 'Common' THEN 0
        WHEN quality = 'Uncommon' THEN 1
        WHEN quality = 'Limited' THEN 2
        WHEN quality = 'Rare' THEN 3
        WHEN quality = 'Very Rare' THEN 4
        WHEN quality = 'Import' THEN 5
        WHEN quality = 'Exotic' THEN 6
        WHEN quality = 'Black Market' THEN 7
        WHEN quality = 'Premium' THEN 8
        WHEN quality = 'Legacy' THEN 9
        ELSE 99
      END desc`
		case "quality_asc":
			sortString = `CASE
        WHEN quality = 'Common' THEN 0
        WHEN quality = 'Uncommon' THEN 1
        WHEN quality = 'Limited' THEN 2
        WHEN quality = 'Rare' THEN 3
        WHEN quality = 'Very Rare' THEN 4
        WHEN quality = 'Import' THEN 5
        WHEN quality = 'Exotic' THEN 6
        WHEN quality = 'Black Market' THEN 7
        WHEN quality = 'Premium' THEN 8
        WHEN quality = 'Legacy' THEN 9
        ELSE 99
      END asc`
		default:
			sortString = ""
		}

		tx = tx.Order(sortString)
	}

	tx.Count(&totalCount)

	pagesize := query.PageSize
	page := query.Page

	pagination.SanitizePages(&page, &pagesize)
	offset, limit := pagination.GetOffsetAndLimit(page, pagesize)

	tx = tx.Limit(limit)
	tx = tx.Offset(offset)

	err := tx.Find(&inventoryItems).Error
	if err != nil {
		return nil, err
	}
	return &GetInventoryResponse{
		InventoryItems: inventoryItems,
		Pagination:     pagination.GetPagination(int64(page), int64(pagesize), totalCount),
	}, nil
}
