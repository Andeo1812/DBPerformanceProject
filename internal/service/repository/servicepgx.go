package repository

import (
	"context"
	"database/sql"
	"db-performance-project/internal/pkg"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg/sqltools"
)

type ServiceRepository interface {
	Clear(ctx context.Context) error
	GetStatus(ctx context.Context) (*models.StatusService, error)
}

type servicePostgres struct {
	database *sqltools.Database
}

func NewServicePostgres(database *sqltools.Database) ServiceRepository {
	return &servicePostgres{
		database,
	}
}

func (s servicePostgres) Clear(ctx context.Context) error {
	errMain := sqltools.RunTxOnConn(ctx, pkg.TxInsertOptions, s.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		tx.QueryRowContext(ctx, clearAllTables)
		// if row.err() != nil {
		//	return errors.WithMessagef(pkg.ErrWorkDatabase,
		//		"Err: params input: query - [%s], values - [%s, %s, %s, %s]. Special error: [%s]",
		//		createUser, user.Nickname, user.FullName, user.About, user.Email, rowUser.Err())
		// }

		return nil
	})

	return errMain
}

func (s servicePostgres) GetStatus(ctx context.Context) (*models.StatusService, error) {
	res := &models.StatusService{}

	errMain := sqltools.RunQuery(ctx, s.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowCounters := conn.QueryRowContext(ctx, getCountForumsPostsThreadsUsers)
		// if rowCounters.err() != nil {
		//	return errors.WithMessagef(pkg.ErrWorkDatabase,
		//		"Err: params input: query - [%s], values - [%s, %s, %s, %s]. Special error: [%s]",
		//		createUser, user.Nickname, user.FullName, user.About, user.Email, rowUser.Err())
		// }

		err := rowCounters.Scan(
			&res.Forum,
			&res.Post,
			&res.Thread,
			&res.User)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return nil, errMain
	}

	return res, nil
}
