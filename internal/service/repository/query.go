package repository

const (
	clearAllTables = `
TRUNCATE TABLE forums, posts, threads, user_forums, users, user_votes CASCADE;`

	getCountForumsPostsThreadsUsers = `
SELECT (SELECT count(*) FROM forums) AS forums,
       (SELECT count(*) FROM posts)  AS posts,
       (SELECT count(*) FROM thread) AS threads,
       (SELECT count(*) FROM users)  AS users`
)
