package items

import "context"

type UseCase interface {
	Get(ctx context.Context, query *GetItemQuery) ([]Item, error)
	GetById(ctx context.Context, id uint) (*Item, error)
}

type ItemsUseCase struct {
	repo Repository
}

func NewItemUseCase(itemsRepo Repository) *ItemsUseCase {
	return &ItemsUseCase{
		repo: itemsRepo,
	}
}

func (u *ItemsUseCase) Get(ctx context.Context, query *GetItemQuery) ([]Item, error) {
	return u.repo.Get(ctx, query)
}

func (u *ItemsUseCase) GetById(ctx context.Context, id uint) (*Item, error) {
	return u.repo.GetById(ctx, id)
}
