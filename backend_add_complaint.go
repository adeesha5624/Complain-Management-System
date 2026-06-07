package main

import (
	"encoding/json"
	"net/http"
)

func handleAddComplaint(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var c Complaint
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token := getTokenFromRequest(r)
		user := getUserFromToken(token)
		if user == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if user.Role == "admin" {
			http.Error(w, "Admins may not submit complaints", http.StatusForbidden)
			return
		}

		// Regular users may only create complaints under their own ID.
		c.CustomerID = user.ID

		if c.CustomerID == 0 {
			http.Error(w, "customer_id is required", http.StatusBadRequest)
			return
		}

		err = db.QueryRow(
			"INSERT INTO complaints (customer_id, title, description, status) VALUES ($1, $2, $3, 'Pending') RETURNING id",
			c.CustomerID, c.Title, c.Description,
		).Scan(&c.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(c)
	}
}
