package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/pkg/sqltools"
)

type ThreadRepository interface {
	// Support
	GetThreadIDByForumAndSlug(ctx context.Context, thread *models.Thread) (models.Thread, error)
	CheckExistThread(ctx context.Context, thread *models.Thread) (bool, error)
	GetThreadForumByID(ctx context.Context, thread *models.Thread) (models.Thread, error)
	GetThreadIDBySlug(ctx context.Context, thread *models.Thread) (models.Thread, error)

	CreateThread(ctx context.Context, thread *models.Thread) (models.Thread, error)
	CreatePostsByID(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]models.Post, error)
	GetDetailsThreadByID(ctx context.Context, thread *models.Thread) (models.Thread, error)
	UpdateThreadByID(ctx context.Context, thread *models.Thread) (models.Thread, error)

	// Posts
	GetPostsByIDFlat(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]models.Post, error)
	GetPostsByIDTree(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]models.Post, error)
	GetPostsByIDParentTree(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]models.Post, error)
}

type threadPostgres struct {
	database *sqltools.Database
}

func NewThreadPostgres(database *sqltools.Database) ThreadRepository {
	return &threadPostgres{
		database,
	}
}

func (t threadPostgres) CheckExistThread(ctx context.Context, thread *models.Thread) (bool, error) {
	res := false

	errMain := sqltools.RunQuery(ctx, t.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowThread := conn.QueryRowContext(ctx, checkExistThreadByID, thread.ID)
		if rowThread.Err() != nil {
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d]. Special error: [%s]",
				checkExistThreadByID, thread.ID, rowThread.Err())
		}

		err := rowThread.Scan(&res)
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

func (t threadPostgres) GetThreadIDBySlug(ctx context.Context, thread *models.Thread) (models.Thread, error) {
	res := models.Thread{}

	errMain := sqltools.RunQuery(ctx, t.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowThread := conn.QueryRowContext(ctx, getThreadIDBySlug, thread.Slug)
		if rowThread.Err() != nil {
			if errors.Is(rowThread.Err(), sql.ErrNoRows) {
				return pkg.ErrSuchThreadNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s]. Special error: [%s]",
				getThreadIDByForumAndSlug, thread.Slug, rowThread.Err())
		}

		err := rowThread.Scan(&res.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchThreadNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s]. Special error: [%s]",
				getThreadIDByForumAndSlug, thread.Slug, err)
		}

		return nil
	})

	if errMain != nil {
		return models.Thread{}, errMain
	}

	return res, nil
}

func (t threadPostgres) GetThreadForumByID(ctx context.Context, thread *models.Thread) (models.Thread, error) {
	res := models.Thread{}

	errMain := sqltools.RunQuery(ctx, t.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowThread := conn.QueryRowContext(ctx, getThreadForumByID, thread.ID)
		if rowThread.Err() != nil {
			if errors.Is(rowThread.Err(), sql.ErrNoRows) {
				return pkg.ErrSuchThreadNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s]. Special error: [%s]",
				getThreadIDByForumAndSlug, thread.Slug, rowThread.Err())
		}

		err := rowThread.Scan(&res.Forum)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchThreadNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s]. Special error: [%s]",
				getThreadIDByForumAndSlug, thread.Slug, err)
		}

		return nil
	})

	if errMain != nil {
		return models.Thread{}, errMain
	}

	return res, nil
}

func (t threadPostgres) GetThreadIDByForumAndSlug(ctx context.Context, thread *models.Thread) (models.Thread, error) {
	res := models.Thread{}

	errMain := sqltools.RunQuery(ctx, t.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowThread := conn.QueryRowContext(ctx, getThreadIDByForumAndSlug, thread.Forum, thread.Slug)
		if rowThread.Err() != nil {
			if errors.Is(rowThread.Err(), sql.ErrNoRows) {
				return pkg.ErrSuchThreadNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s]. Special error: [%s]",
				getThreadIDByForumAndSlug, thread.Slug, rowThread.Err())
		}

		err := rowThread.Scan(&res.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchThreadNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s]. Special error: [%s]",
				getThreadIDByForumAndSlug, thread.Slug, err)
		}

		return nil
	})

	if errMain != nil {
		return models.Thread{}, errMain
	}

	return res, nil
}

func (t threadPostgres) CreateThread(ctx context.Context, thread *models.Thread) (models.Thread, error) {
	if thread.Created == "" {
		thread.Created = time.Now().Format(time.RFC3339)
	}

	errMain := sqltools.RunTxOnConn(ctx, pkg.TxInsertOptions, t.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowThread := tx.QueryRowContext(ctx, createForumThread, thread.Title, thread.Author, thread.Forum, thread.Message, thread.Slug, thread.Created)
		if rowThread.Err() != nil {
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s, %s, %s, %s, %s, %s]. Special error: [%s]",
				createForumThread, thread.Title, thread.Author, thread.Forum, thread.Message, thread.Slug, thread.Created, rowThread.Err())
		}

		err := rowThread.Scan(&thread.ID)
		if err != nil {
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s, %s, %s, %s, %s, %s]. Special error: [%s]",
				createForumThread, thread.Title, thread.Author, thread.Forum, thread.Message, thread.Slug, thread.Created, err)
		}

		return nil
	})

	if errMain != nil {
		return models.Thread{}, errMain
	}

	if thread.Forum == thread.Slug {
		thread.Slug = ""
	}

	return *thread, nil
}

func (t threadPostgres) CreatePostsByID(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]models.Post, error) {
	// Defining sending parameters
	query := insertPosts

	countAttributes := strings.Count(query, ",") + 1

	pos := 0

	countInserts := len(posts)

	values := make([]interface{}, countInserts*countAttributes)

	insertTimeString := time.Now().Format(time.RFC3339)

	for i := 0; i < len(posts); i++ {
		values[pos] = posts[i].Parent
		pos++
		values[pos] = posts[i].Author.Nickname
		pos++
		values[pos] = posts[i].Message
		pos++
		values[pos] = thread.Forum
		pos++
		values[pos] = thread.ID
		pos++
		values[pos] = insertTimeString
		pos++
	}

	insertStatement := sqltools.CreateFullQuery(query, countInserts, countAttributes)

	insertStatement += " RETURNING post_id;"

	rows, err := sqltools.InsertBatch(ctx, t.database.Connection, insertStatement, values)
	if err != nil {
		return nil, err
	}

	res := make([]models.Post, len(posts))

	i := 0
	for rows.Next() {
		err = rows.Scan(&res[i].ID)
		if err != nil {
			return nil, err
		}

		res[i].Created = insertTimeString
		res[i].Parent = posts[i].Parent
		res[i].Author.Nickname = posts[i].Author.Nickname
		res[i].Message = posts[i].Message
		res[i].Forum = thread.Forum
		res[i].Thread = thread.ID

		i++
	}

	return res, nil
}

func (t threadPostgres) GetDetailsThreadByID(ctx context.Context, thread *models.Thread) (models.Thread, error) {
	errMain := sqltools.RunQuery(ctx, t.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowThread := conn.QueryRowContext(ctx, getThreadByID, thread.ID)
		if rowThread.Err() != nil {
			if errors.Is(rowThread.Err(), sql.ErrNoRows) {
				return pkg.ErrSuchThreadNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d]. Special error: [%s]",
				getThreadByID, thread.ID, rowThread.Err())
		}

		err := rowThread.Scan(
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

		return nil
	})

	if errMain != nil {
		return models.Thread{}, errMain
	}

	return *thread, nil
}

func (t threadPostgres) UpdateThreadByID(ctx context.Context, thread *models.Thread) (models.Thread, error) {
	errMain := sqltools.RunTxOnConn(ctx, pkg.TxInsertOptions, t.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		rowThread := tx.QueryRowContext(ctx, updateThreadByID, thread.ID, thread.Title, thread.Message)
		if rowThread.Err() != nil {
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d, %s, %s]. Special error: [%s]",
				getThreadByID, thread.ID, thread.Title, thread.Message, rowThread.Err())
		}

		err := rowThread.Scan(
			&thread.Author,
			&thread.Forum,
			&thread.Votes,
			&thread.Slug,
			&thread.Created)
		if err != nil {
			return err
		}

		return nil
	})

	if errMain != nil {
		return models.Thread{}, errMain
	}

	return *thread, nil
}

func (t threadPostgres) GetPostsByIDFlat(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]models.Post, error) {
	var rows *sql.Rows
	var err error

	query := getPostsByFlatBegin

	var values []interface{}

	switch {
	case params.Since != -1 && params.Desc:
		query += " AND post_id < $2"
	case params.Since != -1 && !params.Desc:
		query += " AND post_id > $2"
	case params.Since != -1:
		query += " AND post_id > $2"
	}

	switch {
	case params.Desc:
		query += " ORDER BY created DESC, post_id DESC"
	case !params.Desc:
		query += " ORDER BY created ASC, post_id"
	default:
		query += " ORDER BY created, post_id"
	}

	query += fmt.Sprintf(" LIMIT NULLIF(%d, 0)", params.Limit)

	if params.Since == -1 {
		values = []interface{}{thread.ID}
	} else {
		values = []interface{}{thread.ID, params.Since}
	}

	res := make([]models.Post, 0)

	err = sqltools.RunQuery(ctx, t.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rows, err = conn.QueryContext(ctx, query, values...)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchPostNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%+v]. Special error: [%s]",
				query, values, err)
		}
		defer rows.Close()

		for rows.Next() {
			post := models.Post{}

			timeTmp := time.Time{}

			err = rows.Scan(
				&post.ID,
				&post.Parent,
				&post.Author,
				&post.Message,
				&post.IsEdited,
				&post.Forum,
				&timeTmp)
			if err != nil {
				return err
			}

			post.Thread = thread.ID

			post.Created = timeTmp.Format(time.RFC3339)

			res = append(res, post)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t threadPostgres) GetPostsByIDTree(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]models.Post, error) {
	var rows *sql.Rows
	var err error

	query := getPostsByTreeBegin

	switch {
	case params.Since != -1 && params.Desc:
		query += " AND path < "
	case params.Since != -1 && !params.Desc:
		query += " AND path > "
	case params.Since != -1:
		query += " AND path > "
	}

	if params.Since != -1 {
		query += fmt.Sprintf(` (SELECT path FROM post WHERE post_id = %d) `, params.Since)
	}

	switch {
	case params.Desc:
		query += " ORDER BY path DESC"
	case !params.Desc:
		query += " ORDER BY path ASC, post_id"
	default:
		query += " ORDER BY path, post_id"
	}

	query += fmt.Sprintf(" LIMIT NULLIF(%d, 0)", params.Limit)

	res := make([]models.Post, 0)

	err = sqltools.RunQuery(ctx, t.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rows, err = conn.QueryContext(ctx, query, thread.ID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchPostNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%+v]. Special error: [%s]",
				query, thread.ID, err)
		}
		defer rows.Close()

		for rows.Next() {
			post := models.Post{}

			timeTmp := time.Time{}

			err = rows.Scan(
				&post.ID,
				&post.Parent,
				&post.Author,
				&post.Message,
				&post.IsEdited,
				&post.Forum,
				&timeTmp)
			if err != nil {
				return err
			}

			post.Thread = thread.ID

			post.Created = timeTmp.Format(time.RFC3339)

			res = append(res, post)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t threadPostgres) GetPostsByIDParentTree(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]models.Post, error) {
	var rows *sql.Rows
	var err error

	query := ""

	var values []interface{}

	if params.Since == -1 {
		if params.Desc {
			query = `
					SELECT post_id, parent, author, message, is_edited, forum, created FROM posts
					WHERE path[1] IN (SELECT post_id FROM posts WHERE thread_id = $1 AND parent = 0 ORDER BY post_id DESC LIMIT $2)
					ORDER BY path[1] DESC, path ASC, post_id ASC;`
		} else {
			query = `
					SELECT post_id, parent, author, message, is_edited, forum, created FROM posts
					WHERE path[1] IN (SELECT post_id FROM posts WHERE thread_id = $1 AND parent = 0 ORDER BY post_id ASC LIMIT $2)
					ORDER BY path ASC, post_id ASC;`
		}

		values = []interface{}{thread.ID, params.Limit}
	} else {
		if params.Desc {
			query = `
					SELECT post_id, parent, author, message, is_edited, forum, created FROM posts
					WHERE path[1] IN (SELECT post_id FROM posts WHERE thread_id = $1 AND parent = 0 AND path[1] <
					(SELECT path[1] FROM posts WHERE post_id = $2) ORDER BY post_id DESC LIMIT $3)
					ORDER BY path[1] DESC, path ASC, post_id ASC;`
		} else {
			query = `
					SELECT post_id, parent, author, message, is_edited, forum, created FROM posts
					WHERE path[1] IN (SELECT post_id FROM posts WHERE thread_id = $1 AND parent = 0 AND path[1] >
					(SELECT path[1] FROM posts WHERE post_id = $2) ORDER BY post_id ASC LIMIT $3) 
					ORDER BY path ASC, post_id ASC;`
		}

		values = []interface{}{thread.ID, params.Since, params.Limit}
	}

	res := make([]models.Post, 0)

	err = sqltools.RunQuery(ctx, t.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rows, err = conn.QueryContext(ctx, query, values...)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchPostNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%+v]. Special error: [%s]",
				query, values, err)
		}
		defer rows.Close()

		for rows.Next() {
			post := models.Post{}

			timeTmp := time.Time{}

			err = rows.Scan(
				&post.ID,
				&post.Parent,
				&post.Author,
				&post.Message,
				&post.IsEdited,
				&post.Forum,
				&timeTmp)
			if err != nil {
				return err
			}

			post.Thread = thread.ID

			post.Created = timeTmp.Format(time.RFC3339)

			res = append(res, post)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
