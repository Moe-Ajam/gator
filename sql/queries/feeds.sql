-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
  )
RETURNING *;

-- name: GetFeeds :many
select * from feeds;


-- name: CreateFeedFollow :one
with inserted_feed_follows as (
  insert into feed_follows (id, created_at, updated_at, user_id, feed_id)
  values (
    $1,
    $2,
    $3,
    $4,
    $5
    ) RETURNING *
)
select inserted_feed_follows.*, feeds.name as feed_name, users.name as user_name
from inserted_feed_follows
inner join feeds on inserted_feed_follows.feed_id = feeds.id
inner join users on inserted_feed_follows.user_id = users.id;

-- name: GetFeedByUrl :one
select * from feeds where  url = $1;

-- name: GetFeedsFollowsForUser :many
select feeds.name as feed_name, users.name as user_name from feed_follows
inner join feeds on feeds.id = feed_follows.feed_id
inner join users on users.id = feed_follows.user_id
where feed_follows.user_id = $1;

-- name: DeleteFollowByUser :exec
delete from feed_follows where user_id = $1 and feed_id = $2;
