package handler

import (
	"app/internal/app/model"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllTasks(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	tasks := []*model.Task{}

	rows, err := db.Query("SELECT id, title, priority, deadline, done, project_id, created_at, updated_at, deleted_at FROM tasks")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		t := &model.Task{}
		if err := rows.Scan(&t.Id, &t.Title, &t.Priority, &t.Deadline, &t.Done, &t.ProjectId, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt); err != nil {
			log.Println(err)
		}
		tasks = append(tasks, t)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
	}
	if len(tasks) < 1 {
		respondJson(w, r, http.StatusOK, map[string]string{"status": "is empty"})
		return
	}

	respondJson(w, r, http.StatusOK, tasks)

}

func GetProjectTasks(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	tasks := []*model.Task{}
	vars := mux.Vars(r)
	title := vars["title"]

	rows, err := db.Query("SELECT t.id, t.title, t.priority, t.deadline, t.done, t.project_id, t.created_at, t.updated_at, t.deleted_at FROM tasks t INNER JOIN projects p ON t.project_id = p.id WHERE p.title=?", title)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		t := &model.Task{}
		if err := rows.Scan(&t.Id, &t.Title, &t.Priority, &t.Deadline, &t.Done, &t.ProjectId, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt); err != nil {
			log.Println(err)
		}
		tasks = append(tasks, t)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
	}

	if len(tasks) < 1 {
		respondError(w, r, http.StatusNotFound, errors.New("Record not found"))
		return
	}

	respondJson(w, r, http.StatusOK, tasks)

}

func CreateTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	t := &model.Task{}
	vars := mux.Vars(r)
	title := vars["title"]

	if err := json.NewDecoder(r.Body).Decode(t); err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}

	ins, err := db.Query("INSERT tasks SET title=?, priority=?, deadline=?, done=?, project_id=(SELECT id FROM projects WHERE title=?), created_at=NOW(), updated_at=NOW(), deleted_at=NOW()",
		t.Title, t.Priority, t.Deadline, t.Done, title)
	if err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}
	defer ins.Close()

	respondJson(w, r, http.StatusCreated, map[string]string{"success": "OK"})
}

func UpdateTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	t := &model.Task{}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := json.NewDecoder(r.Body).Decode(t); err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}

	upd, err := db.Query("UPDATE tasks SET title=?, priority=?, deadline=?, done=?, updated_at=NOW() WHERE id = ?",
		t.Title, t.Priority, t.Deadline, t.Done, id)
	defer upd.Close()
	if err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}

	respondJson(w, r, http.StatusCreated, map[string]string{"success": "OK"})
}

func DeleteTask(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	del, err := db.Query("DELETE FROM tasks WHERE id = ?", id)
	defer del.Close()
	if err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}

	respondJson(w, r, http.StatusOK, map[string]string{"status": "OK"})
}
