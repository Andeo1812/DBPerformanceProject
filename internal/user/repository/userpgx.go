package repository

import (
	"context"
	"database/sql"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/pkg/sqltools"
	"github.com/pkg/errors"
)

type UserRepository interface {
	// Support
	CheckFreeEmail(ctx context.Context, user *models.User) (bool, error)

	CreateUser(ctx context.Context, user *models.User) (models.User, error)
	GetUserByEmailOrNickname(ctx context.Context, user *models.User) ([]models.User, error)
	GetUserByNickname(ctx context.Context, user *models.User) (models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (models.User, error)
}

type userPostgres struct {
	database *sqltools.Database
}

func NewUserPostgres(database *sqltools.Database) UserRepository {
	return &userPostgres{
		database,
	}
}

func (u userPostgres) CheckFreeEmail(ctx context.Context, user *models.User) (bool, error) {
	res := false

	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		row := conn.QueryRowContext(ctx, checkFreeEmail, user.Email)
		if row.Err() != nil {
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s]. Special error: [%s]",
				checkFreeEmail, user.Email, row.Err())
		}

		err := row.Scan(&res)
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

func (u userPostgres) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
	errMain := sqltools.RunTxOnConn(ctx, pkg.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowUser := tx.QueryRowContext(ctx, createUser, user.Nickname, user.FullName, user.About, user.Email)
		if rowUser.Err() != nil {
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s, %s, %s, %s]. Special error: [%s]",
				createUser, user.Nickname, user.FullName, user.About, user.Email, rowUser.Err())
		}

		return nil
	})

	if errMain != nil {
		return models.User{}, errMain
	}

	return *user, nil
}

func (u userPostgres) GetUserByEmailOrNickname(ctx context.Context, user *models.User) ([]models.User, error) {
	res := make([]models.User, 0)

	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowsUsers, err := conn.QueryContext(ctx, getUserByEmailOrNickname, user.Nickname, user.Email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchUserNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s, %s]. Special error: [%s]",
				getUserByEmailOrNickname, user.Nickname, user.Email, err)
		}
		defer rowsUsers.Close()

		for rowsUsers.Next() {
			values := models.User{}

			err = rowsUsers.Scan(
				&values.Nickname,
				&values.FullName,
				&values.About,
				&values.Email)
			if err != nil {
				return err
			}

			res = append(res, values)
		}

		if len(res) == 0 {
			return pkg.ErrSuchUserNotFound
		}

		return nil
	})

	if errMain != nil {
		return []models.User{}, errMain
	}

	return res, nil
}

func (u userPostgres) GetUserByNickname(ctx context.Context, user *models.User) (models.User, error) {
	res := models.User{}

	errMain := sqltools.RunQuery(ctx, u.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowUser := conn.QueryRowContext(ctx, getUserByNickname, user.Nickname)
		if rowUser.Err() != nil {
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s]. Special error: [%s]",
				getUserByNickname, user.Nickname, rowUser.Err())
		}

		err := rowUser.Scan(
			&res.FullName,
			&res.About,
			&res.Email,
			&res.Nickname)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchUserNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s]. Special error: [%s]",
				getUserByNickname, user.Nickname, rowUser.Err())
		}

		return nil
	})

	if errMain != nil {
		return models.User{}, errMain
	}

	return res, nil
}

func (u userPostgres) UpdateUser(ctx context.Context, user *models.User) (models.User, error) {
	res := models.User{}

	errMain := sqltools.RunTxOnConn(ctx, pkg.TxInsertOptions, u.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowUser := tx.QueryRowContext(ctx, updateUser, user.FullName, user.About, user.Email, user.Nickname)
		if rowUser.Err() != nil {
			return errors.WithMessagef(pkg.ErrUpdateUserDataConflict,
				"Err: params input: query - [%s], values - [%s, %s, %s]. Special error: [%s]",
				updateUser, user.FullName, user.About, user.Email, rowUser.Err())
		}

		err := rowUser.Scan(
			&res.FullName,
			&res.About,
			&res.Email,
			&res.Nickname)
		if err != nil {
			return err
		}
		return nil
	})

	if errMain != nil {
		return models.User{}, errMain
	}

	return res, nil
}
