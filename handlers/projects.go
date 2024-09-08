package handlers

import (
	"encoding/json"
	"go-backend-api/database"
	"go-backend-api/models"
	"net/http"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, date FROM project")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		if err := rows.Scan(&project.ID, &project.Name, &project.Date); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		projects = append(projects, project)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
