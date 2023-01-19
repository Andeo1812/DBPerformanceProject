package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/pkg/sqltools"
)

type ForumRepository interface {
	CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error)
	GetDetailsForum(ctx context.Context, forum *models.Forum) (*models.Forum, error)
	GetThreads(ctx context.Context, forum *models.Forum, params *pkg.GetThreadsParams) ([]*models.Thread, error)
	GetUsers(ctx context.Context, forum *models.Forum, params *pkg.GetUsersParams) ([]*models.User, error)
}

type forumPostgres struct {
	database *sqltools.Database
}

func NewForumPostgres(database *sqltools.Database) ForumRepository {
	return &forumPostgres{
		database,
	}
}

func (f forumPostgres) CreateForum(ctx context.Context, forum *models.Forum) (*models.Forum, error) {
	errMain := sqltools.RunTxOnConn(ctx, pkg.TxInsertOptions, f.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, createForum, forum.Title, forum.User, forum.Slug)
		if row.Err() != nil {
			return pkg.ErrUpdateUserDataConflict
		}

		// if row.err() != nil {
		//	return errors.WithMessagef(pkg.ErrWorkDatabase,
		//		"Err: params input: query - [%s], values - [%s, %s, %s, %s]. Special error: [%s]",
		//		createUser, user.Nickname, user.FullName, user.About, user.Email, rowUser.Err())
		// }

		return nil
	})

	return forum, errMain
}

func (f forumPostgres) GetDetailsForum(ctx context.Context, forum *models.Forum) (*models.Forum, error) {
	errMain := sqltools.RunQuery(ctx, f.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowCounters := conn.QueryRowContext(ctx, getForumBySlug, forum.Slug)
		if errors.Is(rowCounters.Err(), sql.ErrNoRows) {
			return pkg.ErrSuchForumNotFound
		}

		// else {
		//	return errors.WithMessagef(pkg.ErrWorkDatabase,
		//		"Err: params input: query - [%s], values - [%s, %s, %s, %s]. Special error: [%s]",
		//		createUser, user.Nickname, user.FullName, user.About, user.Email, rowUser.Err())
		// }

		err := rowCounters.Scan(
			&forum.Title,
			&forum.User,
			&forum.Posts,
			&forum.Threads)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return nil, errMain
	}

	return forum, nil
}

func (f forumPostgres) GetThreads(ctx context.Context, forum *models.Forum, params *pkg.GetThreadsParams) ([]*models.Thread, error) {
	query := getForumThreadsBegin

	orderBy := "ORDER BY t.created "
	querySince := " AND t.created >= $2 "

	var rows *sql.Rows
	var err error

	if params.Desc {
		orderBy += "DESC"
	}

	orderBy += fmt.Sprintf(" LIMIT %d", params.Limit)

	switch {
	case params.Since != "" && params.Desc:
		querySince = " and t.created <= $2 "
	case params.Since != "" && !params.Desc:
		querySince = " and t.created >= $2 "
	}

	var values []interface{}

	if params.Since != "" {
		query += querySince + orderBy

		values = []interface{}{forum.Slug, params.Since}
	} else {
		query += orderBy

		values = []interface{}{forum.Slug}
	}

	res := make([]*models.Thread, 0)

	err = sqltools.RunQuery(ctx, f.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rows, err = conn.QueryContext(ctx, query, values...)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			thread := &models.Thread{}

			err = rows.Scan(
				&thread.ID,
				&thread.Title,
				&thread.Author,
				&thread.Forum,
				&thread.Message,
				&thread.Votes,
				&thread.Slug,
				&thread.Created)
			if err != nil {
				return err
			}

			res = append(res, thread)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (f forumPostgres) GetUsers(ctx context.Context, forum *models.Forum, params *pkg.GetUsersParams) ([]*models.User, error) {
	var rows *sql.Rows
	var err error

	query := getForumUsersBegin

	switch {
	case params.Desc && params.Since != "":
		query += fmt.Sprintf(" and u.nickname < '%s'", params.Since)
	case params.Since != "":
		query += fmt.Sprintf(" and u.nickname > '%s'", params.Since)
	}

	query += " ORDER BY u.nickname "

	if params.Desc {
		query += "DESC"
	}

	query += fmt.Sprintf(" LIMIT %d", params.Limit)

	res := make([]*models.User, 0)

	err = sqltools.RunQuery(ctx, f.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rows, err = conn.QueryContext(ctx, query, forum.Slug)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			user := &models.User{}

			err = rows.Scan(
				&user.Nickname,
				&user.FullName,
				&user.About,
				&user.Email)
			if err != nil {
				return err
			}

			res = append(res, user)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
