package auth

import (
	"context"
	"fmt"
	"github.com/MuhammadIbraAlfathar/go-online-shop/external/database"
	"github.com/MuhammadIbraAlfathar/go-online-shop/internal/config"
	"github.com/google/uuid"
	"github.com/test-go/testify/assert"
	"testing"
)

var s service

func init() {
	filename := "../../cmd/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)

	if err != nil {
		panic(err)
	}

	r := newRepository(db)
	s = newService(r)
}

func TestRegister_Success(t *testing.T) {
	req := RequestRegisterPayload{
		Password: "test12345",
		Email:    fmt.Sprintf("%v@gmail.com", uuid.NewString()),
	}
	err := s.register(context.Background(), req)
	assert.Nil(t, err)
}

func TestEmailAlreadyUse(t *testing.T) {
	req := RequestRegisterPayload{Password: "test12345", Email: "ibraalfathar12@gmail.com"}
	err := s.register(context.Background(), req)
	assert.NotNil(t, err)
}
