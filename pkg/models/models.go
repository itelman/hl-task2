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

func (t *TaskRequest) getMD5Hash() string {
	hash := md5.Sum([]byte(t.Title + t.ActiveAt))
	return hex.EncodeToString(hash[:])
}

func NewTask(request TaskRequest) Task {
	return Task{request.getMD5Hash(), request.Title, request.ActiveAt, false}
}

func UpdatedTask(request TaskRequest, id string) Task {
	return Task{id, request.Title, request.ActiveAt, false}
}

func (t *Task) Check() {
	t.Done = true
}

func (t *Task) MarkAsWeekend() {
	t.Title = "ВЫХОДНОЙ - " + t.Title
}
