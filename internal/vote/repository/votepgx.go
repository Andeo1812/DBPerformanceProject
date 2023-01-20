package repository

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/pkg/sqltools"
)

type VoteRepository interface {
	CheckExistVote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (bool, error)
	UpdateVote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) error
	CreateVote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) error
}

type votePostgres struct {
	database *sqltools.Database
}

func NewVotePostgres(database *sqltools.Database) VoteRepository {
	return &votePostgres{
		database,
	}
}

func (v votePostgres) CheckExistVote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) (bool, error) {
	res := false

	errMain := sqltools.RunQuery(ctx, v.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowExist := conn.QueryRowContext(ctx, checkExists, params.Nickname, thread.ID)
		if rowExist.Err() != nil {
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s, %d]. Special error: [%s]",
				checkExists, params.Nickname, thread.ID, rowExist.Err())
		}

		err := rowExist.Scan(&res)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return false, errMain
	}

	return res, nil
}

func (v votePostgres) UpdateVote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) error {
	errMain := sqltools.RunTxOnConn(ctx, pkg.TxInsertOptions, v.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowUpdate := tx.QueryRowContext(ctx, updateVote, thread.ID, params.Nickname, params.Voice)
		if rowUpdate.Err() != nil {
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s, %d, %d]. Special error: [%s]",
				updateVote, params.Nickname, thread.ID, params.Voice, rowUpdate.Err())
		}

		return nil
	})

	return errMain
}

func (v votePostgres) CreateVote(ctx context.Context, thread *models.Thread, params *pkg.VoteParams) error {
	errMain := sqltools.RunTxOnConn(ctx, pkg.TxInsertOptions, v.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowCreate := tx.QueryRowContext(ctx, createVote, params.Nickname, thread.ID, params.Voice)
		if rowCreate.Err() != nil {
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s, %d, %d]. Special error: [%s]",
				createVote, params.Nickname, thread.ID, params.Voice, rowCreate.Err())
		}

		return nil
	})

	return errMain
}
