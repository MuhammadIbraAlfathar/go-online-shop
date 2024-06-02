package auth

import (
	"github.com/MuhammadIbraAlfathar/go-online-shop/infra/response"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type Role string

const (
	ROLE_Admin Role = "admin"
	ROLE_User  Role = "user"
)

type AuthenticationEntity struct {
	Id        int       `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      Role      `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewFromRegisterRequest(req RequestRegisterPayload) AuthenticationEntity {
	return AuthenticationEntity{
		Email:     req.Email,
		Password:  req.Password,
		Role:      ROLE_User,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (a AuthenticationEntity) Validate() (err error) {
	if err = a.ValidateEmail(); err != nil {
		return
	}

	if err = a.ValidatePassword(); err != nil {
		return
	}
	return
}

func (a AuthenticationEntity) ValidateEmail() (err error) {

	if a.Email == "" {
		return response.ErrEmailRequired
	}

	emails := strings.Split(a.Email, "@")
	if len(emails) != 2 {
		return response.ErrEmailInvalid
	}

	return
}

func (a AuthenticationEntity) ValidatePassword() (err error) {
	if a.Password == "" {
		return response.ErrPasswordRequired
	}

	if len(a.Password) < 6 {
		return response.ErrPasswordInvalidLength
	}

	return
}

func (a *AuthenticationEntity) EncryptPassword(salt int) (err error) {
	passwordEncrypt, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	a.Password = string(passwordEncrypt)
	return nil
}

func (a AuthenticationEntity) IsExits() bool {
	//if a.Id == 0 {
	//	return response.ErrAuthIsNotExists
	//}

	return a.Id != 0
}

func (a AuthenticationEntity) VerifyPasswordFromEncrypted(plain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plain))
}

func (a AuthenticationEntity) VerifyPasswordFromPlain(encrypted string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(a.Password))
}