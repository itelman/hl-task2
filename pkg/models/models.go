package models

import (
	"crypto/md5"
	"encoding/hex"
)

type TaskRequest struct {
	Title    string `json:"title"`
	ActiveAt string `json:"activeAt"`
}

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	ActiveAt string `json:"activeAt"`
	Done     bool
}

func (t *Task) getMD5Hash(request TaskRequest) string {
	hash := md5.Sum([]byte(request.Title))
	return hex.EncodeToString(hash[:])
}

func (t *Task) Set(request TaskRequest) error {
	t.ID = t.getMD5Hash(request)
	t.Title = request.Title
	t.ActiveAt = request.ActiveAt

	return nil
}
