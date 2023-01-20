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
WHERE LOWER(slug) = LOWER($1);`

	checkExistThreadByID = `
SELECT EXISTS(SELECT 1 FROM threads WHERE thread_id = $1);`

	getThreadIDByForumAndSlug = `
SELECT thread_id
FROM threads
WHERE LOWER(forum) = LOWER($1) AND LOWER(slug) = LOWER($2);`

	getThreadForumByID = `
SELECT forum
FROM threads
WHERE thread_id = $1;`

	getThreadIDBySlug = `
SELECT thread_id
FROM threads
WHERE LOWER(slug) = LOWER($1);`

	createForumThread = `
INSERT INTO threads(title, author, forum, message, slug, created)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING thread_id;`

	updateThreadByID = `
UPDATE threads
SET title   = $2,
    message = $3
WHERE thread_id = $1
RETURNING author, forum, votes, slug, created;`

	insertPosts = "INSERT INTO posts(parent, author, message, forum, thread_id, created) VALUES "

	getPostsByFlatBegin = `
SELECT post_id, parent, author, message, is_edited, forum, created FROM posts WHERE thread_id = $1 `

	getPostsByTreeBegin = `
SELECT post_id, parent, author, message, is_edited, forum, created FROM posts WHERE thread_id = $1 `
)
