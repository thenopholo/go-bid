package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thenopholo/go-bid/internal/config"
	"github.com/thenopholo/go-bid/internal/store"
	"github.com/thenopholo/go-bid/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	pool    *pgxpool.Pool
	queries *store.Queries
	logger  *config.Logger
}

func NewUserService(pool *pgxpool.Pool) *UserService {
	logger := config.NewLogger("USER_SERVICE")

	return &UserService{
		pool:    pool,
		queries: store.New(pool),
		logger:  logger,
	}
}

func (us *UserService) CreateUser(ctx context.Context, userName, email, password, bio string) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		us.logger.Errf("error on creating new user: %v", err)
		return uuid.UUID{}, err
	}

	args := store.CreateUserParams{
		Name:         userName,
		Email:        email,
		PasswordHash: hash,
		Bio:          bio,
	}

	user, err := us.queries.CreateUser(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			us.logger.Err("error 23505 on DB")
			return uuid.UUID{}, utils.ErrDuplicateUserNameOrPassword
		}
		us.logger.Errf("error on creating new user: %v", err)
		return uuid.UUID{}, err
	}

	return user.ID, nil
}
