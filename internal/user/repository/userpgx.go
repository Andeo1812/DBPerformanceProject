package repository

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/pkg/sqltools"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmailOrNickname(ctx context.Context, user *models.User) ([]*models.User, error)
	GetUserByNickname(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
}

type userPostgres struct {
	database *sqltools.Database
}

func NewUserPostgres(database *sqltools.Database) UserRepository {
	return &userPostgres{
		database,
	}
}

func (u userPostgres) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	sqltools.RunTxOnConn(ctx, pkg.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowUser := tx.QueryRowContext(ctx, createUser, user.Nickname, user.FullName, user.About, user.Email)
		if errors.Is(rowUser.Err(), sql.ErrTxDone) {
			return pkg.ErrSuchUserExist
		}

		// else {
		//	return errors.WithMessagef(pkg.ErrWorkDatabase,
		//		"Err: params input: query - [%s], values - [%s, %s, %s, %s]. Special error: [%s]",
		//		createUser, user.Nickname, user.FullName, user.About, user.Email, rowUser.Err())
		// }

		return nil
	})

	// if errMain != nil {
	//	return nil, errMain
	// }

	return user, nil
}

func (u userPostgres) GetUserByEmailOrNickname(ctx context.Context, user *models.User) ([]*models.User, error) {
	panic("implement me")
}

func (u userPostgres) GetUserByNickname(ctx context.Context, user *models.User) (*models.User, error) {
	panic("implement me")
}

func (u userPostgres) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	panic("implement me")
}
