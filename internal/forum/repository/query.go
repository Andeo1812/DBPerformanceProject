package repository

const (
	getForumUsersBegin = `
SELECT u.nickname, u.fullname, u.about, u.email
FROM user_forums u
WHERE u.forum = $1 `

	createForum = `
INSERT INTO forums(title, users_nickname, slug)
VALUES ($1, $2, $3);`

	getForumBySlug = `
SELECT title, users_nickname, posts, threads, slug
FROM forums
WHERE slug = $1`

	checkExistForumBySlug = `
SELECT EXISTS(SELECT 1 FROM forums WHERE slug = $1);`

	getForumThreadsBegin = `
SELECT t.thread_id,
       t.title,
       t.author,
       t.forum,
       t.message,
       t.votes,
       t.slug,
       t.created
FROM threads AS t
         LEFT JOIN forums f ON t.forum = f.slug
WHERE f.slug = $1 `
)
