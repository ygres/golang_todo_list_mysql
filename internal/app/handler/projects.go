package handler

import (
	"app/internal/app/model"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllProject(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	p := []*model.Project{}

	rows, err := db.Query("SELECT title, archived, created_at, updated_at, deleted_at FROM projects")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		t := &model.Project{}
		if err := rows.Scan(&t.Title, &t.Archived, &t.CreatedAt, &t.UpdatedAt, &t.DeletedAt); err != nil {
			log.Println(err)
		}
		p = append(p, t)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
	}

	if len(p) < 1 {
		respondJson(w, r, http.StatusOK, map[string]string{"status": "is empty"})
		return
	}

	respondJson(w, r, http.StatusOK, p)
}

func CreateProject(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	p := &model.Project{}

	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}
	p.Unarchive()
	ins, err := db.Query("INSERT projects SET title=?, archived=?, created_at=NOW(), updated_at=NOW(), deleted_at=NOW()", p.Title, p.Archived)
	if err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}
	defer ins.Close()

	respondJson(w, r, http.StatusCreated, map[string]string{"success": "OK"})
}

func GetProject(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	p := &model.Project{}

	if err := db.QueryRow("SELECT title, archived, created_at, updated_at, deleted_at FROM projects WHERE title = ?", title).Scan(
		&p.Title, &p.Archived, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			respondError(w, r, http.StatusBadRequest, errors.New("Record not  found"))
			return
		}
	}

	respondJson(w, r, http.StatusOK, p)

}

func UpdateProject(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	p := &model.Project{}

	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}

	upd, err := db.Query("UPDATE projects SET title=?, archived=?, updated_at=NOW() WHERE title = ?", p.Title, p.Archived, title)
	defer upd.Close()
	if err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}

	respondJson(w, r, http.StatusOK, map[string]string{"status": "OK"})

}

func DeleteProject(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	del, err := db.Query("DELETE FROM projects WHERE title = ?", title)
	defer del.Close()
	if err != nil {
		respondError(w, r, http.StatusBadRequest, err)
		return
	}

	respondJson(w, r, http.StatusOK, map[string]string{"status": "OK"})
}
