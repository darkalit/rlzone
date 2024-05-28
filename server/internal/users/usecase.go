package users

import (
	"context"
	"errors"
	"net/http"

	"gorm.io/gorm"

	"github.com/darkalit/rlzone/server/config"
	"github.com/darkalit/rlzone/server/pkg/auth"
	"github.com/darkalit/rlzone/server/pkg/httpErrors"
)

type UseCase interface {
	Register(ctx context.Context, request *RegisterRequest) (*UserWithTokens, error)
	Login(ctx context.Context, request *LoginRequest) (*UserWithTokens, error)
	RefreshToken(ctx context.Context, refreshToken string) (*UserWithTokens, error)
	Logout(ctx context.Context, refreshToken string) error
	Get(ctx context.Context, query *GetUsersQuery) (*GetResponse, error)
	GetById(ctx context.Context, id uint) (*User, error)
	BlockUser(ctx context.Context, id uint) error
}

type UsersUseCase struct {
	repo Repository
	cfg  *config.Config
}

func NewUserUseCase(usersRepo Repository, cfg *config.Config) *UsersUseCase {
	return &UsersUseCase{
		repo: usersRepo,
		cfg:  cfg,
	}
}

func (u *UsersUseCase) Register(
	ctx context.Context,
	request *RegisterRequest,
) (*UserWithTokens, error) {
	existingUser, err := u.repo.GetByEmail(ctx, request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, httpErrors.NewRestError(http.StatusBadRequest, "Email already exists", nil)
	}

	createdUser := User{
		Email:    request.Email,
		EpicID:   request.EpicID,
		Password: request.Password,
		Role:     RoleUser,
	}
	err = u.repo.Create(ctx, &createdUser)
	if err != nil {
		return nil, err
	}

	payload := auth.JWTPayload{
		UserID: createdUser.ID,
		Role:   string(createdUser.Role),
		Email:  createdUser.Email,
	}
	refreshToken, err := auth.GenJWT(&payload, u.cfg, auth.RefreshTokenType)
	if err != nil {
		return nil, err
	}

	createdToken := Token{
		RefreshToken: refreshToken,
		UserID:       createdUser.ID,
	}
	err = u.repo.CreateToken(ctx, &createdToken)
	if err != nil {
		return nil, err
	}

	return &UserWithTokens{
		User:         createdUser,
		RefreshToken: refreshToken,
	}, nil
}

func (u *UsersUseCase) Login(ctx context.Context, request *LoginRequest) (*UserWithTokens, error) {
	foundUser, err := u.repo.GetByEmail(ctx, request.Email)
	if err != nil {
		return nil, httpErrors.NewRestError(http.StatusUnauthorized, "Wrong login credentials", nil)
	}

	err = foundUser.ComparePasswords(request.Password)
	if err != nil {
		return nil, httpErrors.NewRestError(http.StatusUnauthorized, "Wrong login credentials", nil)
	}

	payload := auth.JWTPayload{
		UserID: foundUser.ID,
		Role:   string(foundUser.Role),
		Email:  foundUser.Email,
	}
	refreshToken, err := auth.GenJWT(&payload, u.cfg, auth.RefreshTokenType)
	if err != nil {
		return nil, err
	}

	return &UserWithTokens{
		User:         *foundUser,
		RefreshToken: refreshToken,
	}, nil
}

func (u *UsersUseCase) RefreshToken(
	ctx context.Context,
	refreshToken string,
) (*UserWithTokens, error) {
	payload, err := auth.VerifyToken(refreshToken, u.cfg, auth.RefreshTokenType)
	if err != nil {
		return nil, httpErrors.NewRestError(http.StatusUnauthorized, "Refresh token invalid", err)
	}

	foundUser, err := u.repo.GetById(ctx, payload.UserID)
	if err != nil {
		return nil, httpErrors.NewRestError(http.StatusNotFound, "User not found", err)
	}

	payload = &auth.JWTPayload{
		UserID: foundUser.ID,
		Role:   string(foundUser.Role),
		Email:  foundUser.Email,
	}
	newRefreshToken, err := auth.GenJWT(payload, u.cfg, auth.RefreshTokenType)
	if err != nil {
		return nil, err
	}

	return &UserWithTokens{
		User:         *foundUser,
		RefreshToken: newRefreshToken,
	}, nil
}

func (u *UsersUseCase) Logout(ctx context.Context, refreshToken string) error {
	payload, err := auth.VerifyToken(refreshToken, u.cfg, auth.RefreshTokenType)
	if err != nil {
		return httpErrors.NewRestError(http.StatusUnauthorized, "Refresh token invalid", err)
	}

	err = u.repo.DeleteTokenByUserId(ctx, payload.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UsersUseCase) BlockUser(ctx context.Context, id uint) error {
	foundUser, err := u.repo.GetById(ctx, id)
	if err != nil {
		return httpErrors.NewRestError(http.StatusNotFound, "User not found", err)
	}

	foundUser.IsBlocked = true

	err = u.repo.Update(ctx, foundUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *UsersUseCase) Get(ctx context.Context, query *GetUsersQuery) (*GetResponse, error) {
	return u.repo.Get(ctx, query)
}

func (u *UsersUseCase) GetById(ctx context.Context, id uint) (*User, error) {
	return u.repo.GetById(ctx, id)
}
