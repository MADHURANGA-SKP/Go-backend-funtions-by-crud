-- name: Createtasks :one
INSERT INTO tasks(
    taskid,
    taksname,
    tasktime,
    taskdate
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: Gettasks :one
SELECT * FROM tasks
WHERE taskid = $1 LIMIT 1;

-- name: Listtasks :many
SELECT * FROM tasks
ORDER BY taskid
LIMIT 1;

-- name: Updatetasks :one
UPDATE tasks SET taskid = $2, taksname = $3, tasktime =$4, taskdate = $5
WHERE taskid = $1
RETURNING *;

-- name: Deletetasks :exec
DELETE FROM tasks 
WHERE taskid = $1;