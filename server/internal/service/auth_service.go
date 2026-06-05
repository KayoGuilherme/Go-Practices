package service

import (
	"context"
	"database/sql"
	"errors"
	"server/internal/domain"
	"server/internal/domain/exceptions/users"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = "secretKey"


type authService struct {
	repo domain.Repository
	timeout time.Duration
}

func NewAuthService(repo domain.Repository) domain.AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Login(ctx context.Context, email, password string) (string, *domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil, users.NewUserInvalidCredentialsError()
		}
		return "", nil, err
	}

	if err := util.CheckPassword(password, user.Password); err != nil {
		return "", nil, users.NewUserInvalidCredentialsError()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtClaims{
		ID:   strconv.FormatInt(user.ID, 10),
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    strconv.FormatInt(user.ID, 10),
			Subject:   user.Email,
		},
	})

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", nil, err
	}

	return signedToken, user, nil
}
