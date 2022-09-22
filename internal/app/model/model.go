package model

import "time"

type Project struct {
	Title     string    `json: "title"`
	Archived  bool      `json: "archived"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	DeletedAt time.Time `json: "Deleted_at"`
}

type Task struct {
	Id        uint      `json: id`
	Title     string    `json: "title"`
	Priority  uint      `json: "priority"`
	Deadline  time.Time `json: "deadline"`
	Done      bool      `json: "done"`
	ProjectId uint      `json: "project_id"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	DeletedAt time.Time `json: "Deleted_at"`
}

func (p *Project) Archive() {
	p.Archived = true
}

func (p *Project) Unarchive() {
	p.Archived = false
}
