package repository

const (
	createUser = `
INSERT INTO users(nickname, fullname, about, email)
VALUES ($1, $2, $3, $4);`

	getUserByEmailOrNickname = `
SELECT nickname, fullname, about, email
FROM users
WHERE nickname = $1
   OR email = $2;`

	getUserByNickname = `
SELECT fullname, about, email
FROM users
where nickname = $1;`

	updateUser = `
UPDATE users
SET fullname = COALESCE(NULLIF(TRIM($1), ''), fullname),
    about    = COALESCE(NULLIF(TRIM($2), ''), about),
    email    = COALESCE(NULLIF(TRIM($3), ''), email)
where nickname = $4
RETURNING fullname, about, email;`
)
