package repository

const (
	createVote = `
INSERT INTO user_votes(nickname, thread_id, voice)
VALUES ($1, $2, $3);`

	updateVote = `
UPDATE user_votes
SET voice = $3
WHERE thread_id = $1
  AND nickname = $2
  AND voice != $3;`

	checkExists = `SELECT EXISTS(SELECT 1 FROM user_votes WHERE nickname = $1 AND thread_id = $2);`
)
