package repository

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"db-performance-project/internal/models"
	"db-performance-project/internal/pkg"
	"db-performance-project/internal/pkg/sqltools"
)

type ThreadRepository interface {
	// Support
	GetThreadIDBySlug(ctx context.Context, thread *models.Thread) (*models.Thread, error)

	CreateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error)
	CreatePostsByID(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]*models.Post, error)
	GetDetailsThreadByID(ctx context.Context, thread *models.Thread) (*models.Thread, error)
	UpdateThreadByID(ctx context.Context, thread *models.Thread) (*models.Thread, error)

	// Posts
	GetPostsByIDFlat(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error)
	GetPostsByIDTree(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error)
	GetPostsByIDParentTree(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error)
}

type threadPostgres struct {
	database *sqltools.Database
}

func NewThreadPostgres(database *sqltools.Database) ThreadRepository {
	return &threadPostgres{
		database,
	}
}

func (t threadPostgres) GetThreadIDBySlug(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	res := &models.Thread{}

	errMain := sqltools.RunQuery(ctx, t.database.Connection, func(ctx context.Context, conn *sql.Conn) error {
		rowThread := conn.QueryRowContext(ctx, getThreadIDBySlug, thread.Slug)
		if errors.As(rowThread.Err(), sql.ErrNoRows) {
			return pkg.ErrSuchThreadNotFound
		}
		// if rowCounters.err() != nil {
		//	return errors.WithMessagef(pkg.ErrWorkDatabase,
		//		"Err: params input: query - [%s], values - [%s, %s, %s, %s]. Special error: [%s]",
		//		createUser, user.Nickname, user.FullName, user.About, user.Email, rowUser.Err())
		// }

		err := rowThread.Scan(&res.ID)
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
func (t threadPostgres) CreateThread(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}

func (t threadPostgres) CreatePostsByID(ctx context.Context, thread *models.Thread, posts []*models.Post) ([]*models.Post, error) {
	panic("implement me")
}

func (t threadPostgres) GetDetailsThreadByID(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}

func (t threadPostgres) UpdateThreadByID(ctx context.Context, thread *models.Thread) (*models.Thread, error) {
	panic("implement me")
}

func (t threadPostgres) GetPostsByIDFlat(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error) {
	panic("implement me")
}

func (t threadPostgres) GetPostsByIDTree(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error) {
	panic("implement me")
}

func (t threadPostgres) GetPostsByIDParentTree(ctx context.Context, thread *models.Thread, params *pkg.GetPostsParams) ([]*models.Post, error) {
	panic("implement me")
}
