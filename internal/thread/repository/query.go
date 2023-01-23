package repository

const (
	getThreadByID = `
SELECT title,
       author,
       forum,
       message,
       votes,
       slug,
       created
FROM threads
WHERE thread_id = $1;`

	getThreadBySlug = `
SELECT thread_id,
       title,
       author,
       forum,
       message,
       votes,
       slug,
       created
FROM threads
WHERE slug = $1;`

	createForumThread = `
INSERT INTO threads(title, author, forum, message, slug, created)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING thread_id;`

	updateThreadByID = `
UPDATE threads
SET title   = COALESCE(NULLIF(TRIM($2), ''), title),
    message = COALESCE(NULLIF(TRIM($3), ''), message)
WHERE thread_id = $1
RETURNING author, forum, votes, slug, created, title, message;`

	insertPosts = "INSERT INTO posts(parent, author, message, forum, thread_id, created) VALUES "

	getPostsByFlatBegin = `
SELECT post_id, parent, author, message, is_edited, forum, created FROM posts WHERE thread_id = $1 `

	getPostsByTreeBegin = `
SELECT post_id, parent, author, message, is_edited, forum, created FROM posts WHERE thread_id = $1 `
)
