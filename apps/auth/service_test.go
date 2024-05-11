package auth

import (
	"context"
	"github.com/MuhammadIbraAlfathar/go-online-shop/external/database"
	"github.com/MuhammadIbraAlfathar/go-online-shop/internal/config"
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
	req := RequestRegisterPayload{Password: "test12345", Email: "ibraalfathar@gmail.com"}
	err := s.register(context.Background(), req)
	assert.Nil(t, err)
}
