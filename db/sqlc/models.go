// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
)

type Task struct {
	Taskid   int32          `json:"taskid"`
	Taksname sql.NullString `json:"taksname"`
	Tasktime sql.NullString `json:"tasktime"`
	Taskdate sql.NullString `json:"taskdate"`
}
