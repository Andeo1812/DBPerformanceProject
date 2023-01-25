package repository

const (
	getPostParent = `
SELECT thread_id
from posts
where post_id = $1;`

	updatePost = `
UPDATE posts
SET message   = COALESCE(NULLIF(TRIM($2), ''), message),
    is_edited = CASE
                    WHEN TRIM($2) = message THEN is_edited
                    ELSE true
        END
WHERE post_id = $1
RETURNING parent, author, forum, thread_id, created, message, is_edited;`

	getPost = `
SELECT parent,
       author,
       message,
       is_edited,
       forum,
       thread_id,
       created
from posts
WHERE post_id = $1;`

	getPostAuthor = `
SELECT nickname,
       fullname,
       about,
       email
FROM users 
WHERE nickname = $1;`

	getPostThread = `
SELECT thread_id,
       title,
       author,
       forum,
       message,
       votes,
       slug,
       created
FROM threads
WHERE thread_id = $1;`

	getPostForum = `
SELECT title,
       users_nickname,
       slug,
       posts,
       threads
FROM forums 
WHERE slug = $1;`
)
