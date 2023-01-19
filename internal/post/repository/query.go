package repository

const (
	checkExistPost = `
SELECT EXISTS(SELECT 1 FROM posts WHERE post_id = $1);`

	getPostParent = `
SELECT thread_id
from posts
where post_id = $1;`

	updatePost = `
UPDATE posts
SET message   = $2,
    is_edited = true
WHERE post_id = $1
RETURNING parent, author, forum, thread_id, created;`

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
SELECT u.nickname,
       u.fullname,
       u.about,
       u.email
from posts AS p
         JOIN users u on u.nickname = p.author
WHERE p.post_id = $1;`

	getPostThread = `
SELECT th.thread_id,
       th.title,
       th.author,
       th.forum,
       th.message,
       th.votes,
       th.slug,
       th.created
from posts AS p
         JOIN threads th on th.id = p.thread_id
WHERE p.post_id = $1;`

	getPostForum = `
SELECT f.title,
       f.users_nickname,
       f.slug,
       f.posts,
       f.threads
from posts AS p
         JOIN forum f on f.slug = p.forum
WHERE p.post_id = $1;`
)
