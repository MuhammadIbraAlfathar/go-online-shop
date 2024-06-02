package auth

import (
	"github.com/MuhammadIbraAlfathar/go-online-shop/infra/response"
	"github.com/test-go/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

func TestValidateEntityAuth(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthenticationEntity{
			Email:    "ibraalfathar@gmail.com",
			Password: "test12345",
		}

		err := authEntity.Validate()
		assert.Nil(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		authEntity := AuthenticationEntity{
			Email:    "ibraalafathar.com",
			Password: "12345678",
		}

		err := authEntity.Validate()
		assert.NotNil(t, err)
	})

	t.Run("fail password length", func(t *testing.T) {
		authEntity := AuthenticationEntity{
			Email:    "ibraalfathar@gmail.com",
			Password: "178",
		}

		err := authEntity.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrPasswordInvalidLength, err)
	})

	t.Run("password is required", func(t *testing.T) {
		authEntity := AuthenticationEntity{
			Email:    "ibraalfathar@gmail.com",
			Password: "",
		}

		err := authEntity.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrPasswordRequired, err)
	})

	t.Run("email is required", func(t *testing.T) {
		authEntity := AuthenticationEntity{
			Email:    "",
			Password: "12345678",
		}

		err := authEntity.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, response.ErrEmailRequired, err)
	})
}

func TestEncryptPassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthenticationEntity{
			Email:    "testing@gmail.com",
			Password: "test12344",
		}

		err := authEntity.EncryptPassword(bcrypt.DefaultCost)

		assert.Nil(t, err)

		log.Printf("%+v\n", authEntity)

	})
}
