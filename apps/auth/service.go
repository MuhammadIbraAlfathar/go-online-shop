package auth

import (
	"context"
	"github.com/MuhammadIbraAlfathar/go-online-shop/infra/response"
	"github.com/MuhammadIbraAlfathar/go-online-shop/internal/config"
)

type Repository interface {
	GetAuthByEmail(ctx context.Context, email string) (model AuthenticationEntity, err error)
	CreateAuth(ctx context.Context, model AuthenticationEntity) (err error)
}

type service struct {
	repo Repository
}

func newService(r Repository) service {
	return service{
		repo: r,
	}
}

func (s service) register(ctx context.Context, req RequestRegisterPayload) (err error) {
	authEntity := NewFromRegisterRequest(req)

	if err = authEntity.Validate(); err != nil {
		return
	}

	if err = authEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
		return
	}

	model, err := s.repo.GetAuthByEmail(ctx, authEntity.Email)
	if err != nil {
		if err != response.ErrNotFound {
			return
		}
	}

	if model.IsExits() {
		return response.ErrEmailAlreadyUse
	}

	return s.repo.CreateAuth(ctx, authEntity)
}
