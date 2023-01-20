package repository

const (
	checkFreeEmail = `
SELECT EXISTS(SELECT 1 FROM users WHERE LOWER(email) = LOWER($1));`

	createUser = `
INSERT INTO users(nickname, fullname, about, email)
VALUES ($1, $2, $3, $4);`

	getUserByEmailOrNickname = `
SELECT nickname, fullname, about, email
FROM users
WHERE LOWER(nickname) = LOWER($1)
   OR LOWER(email) = LOWER($2);`

	getUserByNickname = `
SELECT fullname, about, email, nickname
FROM users
WHERE LOWER(nickname) = LOWER($1);`

	updateUser = `
UPDATE users
SET fullname = COALESCE(NULLIF(TRIM($1), ''), fullname),
    about    = COALESCE(NULLIF(TRIM($2), ''), about),
    email    = COALESCE(NULLIF(TRIM($3), ''), email)
WHERE LOWER(nickname) = LOWER($4) RETURNING fullname, about, email, nickname;`
)
