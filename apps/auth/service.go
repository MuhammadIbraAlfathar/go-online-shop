package auth

import "context"

type Repository interface {
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
	return s.repo.CreateAuth(ctx, authEntity)
}
