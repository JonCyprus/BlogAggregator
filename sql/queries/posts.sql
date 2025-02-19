-- name: CreatePost :one
INSERT INTO posts (created_at, updated_at, url, description, published_at, feed_id)
VALUES (
        NOW(),
        NOW(),
        $1,
        $2,
        $3,
        $4
       ) RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.* FROM posts
INNER JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
INNER JOIN users ON feed_follows.user_id = users.id
WHERE users.id = $1
ORDER BY posts.created_at DESC
LIMIT $2;
