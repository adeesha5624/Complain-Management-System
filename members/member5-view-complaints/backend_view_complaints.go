package main

import (
	"encoding/json"
	"net/http"
)

func handleViewComplaints(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		rows, err := db.Query("SELECT id, customer_id, title, description, status, created_at FROM complaints")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		complaints := []Complaint{}
		for rows.Next() {
			var c Complaint
			err := rows.Scan(&c.ID, &c.CustomerID, &c.Title, &c.Description, &c.Status, &c.CreatedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			complaints = append(complaints, c)
		}

		json.NewEncoder(w).Encode(complaints)
	}
}
