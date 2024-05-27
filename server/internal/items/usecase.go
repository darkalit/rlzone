package items

import "context"

type UseCase interface {
	Get(ctx context.Context, query *GetItemsQuery) (*GetResponse, error)
	GetById(ctx context.Context, id uint) (*Item, error)
	CreateStock(ctx context.Context, createStockRequest *CreateStockRequest) (*Stock, error)
	BuyItem(ctx context.Context, itemId uint, userId uint) (*InventoryItem, error)
	SellItem(ctx context.Context, itemId uint, userId uint) (*InventoryItem, error)
	GetInventory(
		ctx context.Context,
		query *GetItemsQuery,
		userId uint,
	) (*GetInventoryResponse, error)
}

type ItemsUseCase struct {
	repo Repository
}

func NewItemUseCase(itemsRepo Repository) *ItemsUseCase {
	return &ItemsUseCase{
		repo: itemsRepo,
	}
}

func (u *ItemsUseCase) Get(ctx context.Context, query *GetItemsQuery) (*GetResponse, error) {
	return u.repo.Get(ctx, query)
}

func (u *ItemsUseCase) GetById(ctx context.Context, id uint) (*Item, error) {
	return u.repo.GetById(ctx, id)
}

func (u *ItemsUseCase) CreateStock(
	ctx context.Context,
	createStockRequest *CreateStockRequest,
) (*Stock, error) {
	createdStock := Stock{
		Price:       createStockRequest.Price,
		Description: createStockRequest.Description,
		ItemID:      createStockRequest.ItemID,
	}

	err := u.repo.CreateStock(ctx, &createdStock)
	if err != nil {
		return nil, err
	}

	return &createdStock, nil
}

func (u *ItemsUseCase) BuyItem(
	ctx context.Context,
	itemId uint,
	userId uint,
) (*InventoryItem, error) {
	return u.repo.BuyItem(ctx, itemId, userId)
}

func (u *ItemsUseCase) SellItem(
	ctx context.Context,
	itemId uint,
	userId uint,
) (*InventoryItem, error) {
	return u.repo.SellItem(ctx, itemId, userId)
}

func (u *ItemsUseCase) GetInventory(
	ctx context.Context,
	query *GetItemsQuery,
	userId uint,
) (*GetInventoryResponse, error) {
	return u.repo.GetInventory(ctx, query, userId)
}
