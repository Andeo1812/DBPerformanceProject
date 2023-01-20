package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/pkg/sqltools"
)

type PostRepository interface {
	// Support
	GetParentPost(ctx context.Context, post *models.Post) (*models.Post, error)

	UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	GetDetailsPost(ctx context.Context, post *models.Post, params *pkg.PostDetailsParams) (*models.PostDetails, error)
}

type postPostgres struct {
	database *sqltools.Database
}

func NewPostPostgres(database *sqltools.Database) PostRepository {
	return &postPostgres{
		database,
	}
}

func (p postPostgres) GetParentPost(ctx context.Context, post *models.Post) (*models.Post, error) {
	res := &models.Post{}

	errMain := sqltools.RunQuery(ctx, p.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		row := conn.QueryRowContext(ctx, getPostParent, post.Parent)
		if row.Err() != nil {
			if errors.Is(row.Err(), sql.ErrNoRows) {
				return pkg.ErrPostParentNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d]. Special error: [%s]",
				getPostParent, post.Parent, row.Err())
		}

		err := row.Scan(&res.Thread)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrPostParentNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d]. Special error: [%s]",
				getPostParent, post.Parent, err)
		}

		return nil
	})

	if errMain != nil {
		return nil, errMain
	}

	return res, nil
}

func (p postPostgres) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	res := &models.Post{}

	errMain := sqltools.RunTxOnConn(ctx, pkg.TxInsertOptions, p.database.Connection, func(ctx context.Context, tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, updatePost, post.ID, post.Message)
		if row.Err() != nil {
			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s]. Special error: [%s]",
				updatePost, post.Message, row.Err())
		}

		postTime := time.Time{}

		err := row.Scan(
			&res.Parent,
			&res.Author.Nickname,
			&res.Forum,
			&res.Thread,
			&postTime,
			&res.Message,
			&res.IsEdited)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchPostNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%s]. Special error: [%s]",
				updatePost, post.Message, err)
		}

		res.Created = postTime.Format(time.RFC3339)

		res.ID = post.ID

		return nil
	})
	if errMain != nil {
		return nil, errMain
	}

	return res, nil
}

func (p postPostgres) GetDetailsPost(ctx context.Context, post *models.Post, params *pkg.PostDetailsParams) (*models.PostDetails, error) {
	res := &models.PostDetails{}

	res.Post.ID = post.ID

	errMain := sqltools.RunQuery(ctx, p.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		row := conn.QueryRowContext(ctx, getPost, post.ID)
		if row.Err() != nil {
			if errors.Is(row.Err(), sql.ErrNoRows) {
				return pkg.ErrSuchPostNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d]. Special error: [%s]",
				getPost, post.ID, row.Err())
		}

		err := row.Scan(
			&res.Post.Parent,
			&res.Post.Author.Nickname,
			&res.Post.Message,
			&res.Post.IsEdited,
			&res.Post.Forum,
			&res.Post.Thread,
			&res.Post.Created)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return pkg.ErrSuchPostNotFound
			}

			return errors.WithMessagef(pkg.ErrWorkDatabase,
				"Err: params input: query - [%s], values - [%d]. Special error: [%s]",
				getPost, post.ID, err)
		}

		return nil
	})

	if errMain != nil {
		return nil, errMain
	}

	for _, value := range params.Related {
		switch value {
		case pkg.PostDetailForum:
			sqltools.RunQuery(ctx, p.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
				row := conn.QueryRowContext(ctx, getPostForum, post.ID)
				if row.Err() != nil {
					if errors.Is(row.Err(), sql.ErrNoRows) {
						return pkg.ErrSuchPostNotFound
					}

					return errors.WithMessagef(pkg.ErrWorkDatabase,
						"Err: params input: query - [%s], values - [%d]. Special error: [%s]",
						getPostForum, post.ID, row.Err())
				}

				err := row.Scan(
					&res.Forum.Title,
					&res.Forum.User,
					&res.Forum.Slug,
					&res.Forum.Posts,
					&res.Forum.Threads)
				if err != nil {
					if errors.Is(err, sql.ErrNoRows) {
						return pkg.ErrSuchPostNotFound
					}

					return errors.WithMessagef(pkg.ErrWorkDatabase,
						"Err: params input: query - [%s], values - [%d]. Special error: [%s]",
						getPostForum, post.ID, err)
				}

				return nil
			})
		case pkg.PostDetailAuthor:
			sqltools.RunQuery(ctx, p.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
				row := conn.QueryRowContext(ctx, getPostAuthor, post.ID)
				if row.Err() != nil {
					if errors.Is(row.Err(), sql.ErrNoRows) {
						return pkg.ErrSuchPostNotFound
					}

					return errors.WithMessagef(pkg.ErrWorkDatabase,
						"Err: params input: query - [%s], values - [%d]. Special error: [%s]",
						getPostAuthor, post.ID, row.Err())
				}

				err := row.Scan(
					&res.Author.Nickname,
					&res.Author.FullName,
					&res.Author.About,
					&res.Author.Email)
				if err != nil {
					if errors.Is(err, sql.ErrNoRows) {
						return pkg.ErrSuchPostNotFound
					}

					return errors.WithMessagef(pkg.ErrWorkDatabase,
						"Err: params input: query - [%s], values - [%d]. Special error: [%s]",
						getPostAuthor, post.ID, err)
				}

				return nil
			})
		case pkg.PostDetailThread:
			sqltools.RunQuery(ctx, p.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
				row := conn.QueryRowContext(ctx, getPostThread, post.ID)
				if row.Err() != nil {
					if errors.Is(row.Err(), sql.ErrNoRows) {
						return pkg.ErrSuchPostNotFound
					}

					return errors.WithMessagef(pkg.ErrWorkDatabase,
						"Err: params input: query - [%s], values - [%d]. Special error: [%s]",
						getPostThread, post.ID, row.Err())
				}

				err := row.Scan(
					&res.Thread.ID,
					&res.Thread.Title,
					&res.Thread.Author,
					&res.Thread.Forum,
					&res.Thread.Message,
					&res.Thread.Votes,
					&res.Thread.Slug,
					&res.Thread.Created)
				if err != nil {
					if errors.Is(err, sql.ErrNoRows) {
						return pkg.ErrSuchPostNotFound
					}

					return errors.WithMessagef(pkg.ErrWorkDatabase,
						"Err: params input: query - [%s], values - [%d]. Special error: [%s]",
						getPostThread, post.ID, err)
				}

				return nil
			})
		}
	}

	return res, nil
}
