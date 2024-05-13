package users

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/darkalit/rlzone/server/pkg/pagination"
)

type Repository interface {
	CreateToken(ctx context.Context, token *Token) error
	UpdateToken(ctx context.Context, token *Token) error
	CreateOrUpdate(ctx context.Context, token *Token) error
	GetTokenByUserId(ctx context.Context, userId uint) (*Token, error)
	DeleteTokenByUserId(ctx context.Context, userId uint) error

	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Get(ctx context.Context, query *GetUsersQuery) (*GetResponse, error)
	GetById(ctx context.Context, id uint) (*User, error)
	GetByEpicId(ctx context.Context, epicId string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type UsersRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) CreateToken(ctx context.Context, token *Token) error {
	tx := r.db.WithContext(ctx)

	return tx.Create(token).Error
}

func (r *UsersRepo) UpdateToken(ctx context.Context, token *Token) error {
	tx := r.db.WithContext(ctx)

	return tx.Save(token).Error
}

func (r *UsersRepo) CreateOrUpdate(ctx context.Context, token *Token) error {
	tx := r.db.WithContext(ctx)

	err := tx.First(token, "user_id = ?", token.UserID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tx.Create(token).Error
		}

		return err
	}

	return tx.Save(token).Error
}

func (r *UsersRepo) GetTokenByUserId(ctx context.Context, userId uint) (*Token, error) {
	var token Token
	tx := r.db.WithContext(ctx)

	err := tx.First(&token, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *UsersRepo) DeleteTokenByUserId(ctx context.Context, userId uint) error {
	tx := r.db.WithContext(ctx)

	return tx.Delete(&Token{}, "user_id = ?", userId).Error
}

func (r *UsersRepo) Create(ctx context.Context, user *User) error {
	tx := r.db.WithContext(ctx)

	return tx.Create(user).Error
}

func (r *UsersRepo) Update(ctx context.Context, user *User) error {
	tx := r.db.WithContext(ctx)

	return tx.Save(&user).Error
}

func (r *UsersRepo) Get(ctx context.Context, query *GetUsersQuery) (*GetResponse, error) {
	var users []User
	var totalCount int64
	tx := r.db.WithContext(ctx)
	tx = tx.Model(&User{})

	if id := query.ID; id != 0 {
		tx = tx.Where("id = ?", id)
	}

	if epicId := query.EpicID; epicId != "" {
		epicId = "%" + epicId + "%"
		tx = tx.Where("epic_id LIKE ?", epicId)
	}

	if email := query.Email; email != "" {
		tx = tx.Where("email = ?", email)
	}

	if sort := query.Sort; sort != "" {
		var sortString string
		switch sort {
		case "id_desc":
			sortString = "id desc"
		case "id_asc":
			sortString = "id asc"
		case "epic_id_desc":
			sortString = "epic_id desc"
		case "epic_id_asc":
			sortString = "epic_id asc"
		case "email_desc":
			sortString = "email desc"
		case "email_asc":
			sortString = "email asc"
		case "balance_desc":
			sortString = "balance desc"
		case "balance_asc":
			sortString = "balance asc"
		case "creation_desc":
			sortString = "created_at desc"
		case "creation_asc":
			sortString = "created_at asc"
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

	err := tx.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &GetResponse{
		Users:      users,
		Pagination: pagination.GetPagination(int64(page), int64(pagesize), totalCount),
	}, nil
}

func (r *UsersRepo) GetById(ctx context.Context, id uint) (*User, error) {
	var user User

	tx := r.db.WithContext(ctx)
	err := tx.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UsersRepo) GetByEpicId(ctx context.Context, epicId string) (*User, error) {
	var user User

	tx := r.db.WithContext(ctx)
	epicId = "%" + epicId + "%"
	err := tx.Where("epic_id LIKE ?", epicId).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UsersRepo) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	tx := r.db.WithContext(ctx)
	err := tx.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
