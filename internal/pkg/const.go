package pkg

import "database/sql"

// Delivery
const (
	// Validation HTTP
	ContentTypeJSON = "application/json"

	// Validattion size Requests
	BufSizeRequest = 1024 * 1024 * 1
)

type ContextKeyType string

// SessionKey for ctx in auth logic
var SessionKey ContextKeyType = "cookie"

const RequestID = "req-id"

// RequestIDKey for ctx in global middleware
var RequestIDKey ContextKeyType = RequestID

// LoggerKey for ctx in global middleware
var LoggerKey ContextKeyType = "logger"

// Repository

// TxDefaultOptions for Postgres
var TxDefaultOptions = &sql.TxOptions{
	Isolation: sql.LevelDefault,
	ReadOnly:  true,
}

// TxInsertOptions for Postgres
var TxInsertOptions = &sql.TxOptions{
	Isolation: sql.LevelDefault,
	ReadOnly:  false,
}

const (
	TypeSortFlat       = "flat"
	TypeSortTree       = "tree"
	TypeSortParentTree = "parent_tree"

	PostDetailForum  = "forum"
	PostDetailThread = "thread"
	PostDetailAuthor = "user"
)
