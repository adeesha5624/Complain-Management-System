package main

import (
	"encoding/json"
	"net/http"
)

func handleAdminManageComplaint(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if r.Method == "PUT" {
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		// Check authentication and admin role
		token := getTokenFromRequest(r)
		user := getUserFromToken(token)
		if user == nil || user.Role != "admin" {
			http.Error(w, "Unauthorized - Admin access required", http.StatusUnauthorized)
			return
		}

		var updateData map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&updateData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update the complaint with new status
		status, ok := updateData["status"].(string)
		if !ok {
			http.Error(w, "Status field is required", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("UPDATE complaints SET status = $1 WHERE id = $2", status, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Fetch and return the updated complaint
		var c Complaint
		err = db.QueryRow("SELECT id, customer_id, title, description, status, created_at FROM complaints WHERE id = $1", id).Scan(
			&c.ID, &c.CustomerID, &c.Title, &c.Description, &c.Status, &c.CreatedAt,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(c)

	} else if r.Method == "DELETE" {
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		// Check authentication and admin role
		token := getTokenFromRequest(r)
		user := getUserFromToken(token)
		if user == nil || user.Role != "admin" {
			http.Error(w, "Unauthorized - Admin access required", http.StatusUnauthorized)
			return
		}

		result, err := db.Exec("DELETE FROM complaints WHERE id = $1", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if rowsAffected == 0 {
			http.Error(w, "Complaint not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Complaint deleted successfully"})
	}
}
