# Member 5: View Complaints

## 📌 Module Overview
This module acts as the central data feed for the application. It fetches all records from the database and renders them gracefully in a table UI.

## 🔗 Connections & Dependencies
- **API Route**: Hits `GET http://localhost:8080/api/complaints`.
- **Database**: Runs a `SELECT` query on the DB table created by Member 1.
- **Exposed Functionality**: This module exposes `fetchComplaints()` which Member 3 (Add) and Member 4 (Manage) both call to instantly update the UI table when data changes.

## 💻 Core Code Explanation

### 1. Backend Data Retrieval (`backend_view_complaints.go`)
Performs a `SELECT` query, loops through all database rows, scans them into a Go struct array, and returns it as a JSON payload.
```go
rows, err := db.Query("SELECT id, customer_id, title, description, status, created_at FROM complaints")
complaints := []Complaint{}
for rows.Next() {
    var c Complaint
    rows.Scan(&c.ID, &c.CustomerID, &c.Title, &c.Description, &c.Status, &c.CreatedAt)
    complaints = append(complaints, c)
}
json.NewEncoder(w).Encode(complaints)
```

### 2. Frontend List Rendering (`frontend_view_complaints.vue`)
Fetches the JSON array and uses Vue's `v-for` directive to dynamically render table rows.
```html
<tr v-for="item in complaints" :key="item.id" class="table-row">
    <td class="id-col">#{{ item.id }}</td>
    <td><strong>{{ item.customer_id }}</strong></td>
    <td>
        <div class="title-cell">{{ item.title }}</div>
        <div class="desc-cell">{{ item.description }}</div>
    </td>
    <td>
        <span class="status-badge" :class="statusClass(item.status)">{{ item.status }}</span>
    </td>
</tr>
```

## 🚀 Viva Presentation Notes
Focus on how Vue reactivity works here. Explain that the `complaints` array is bound to the template, so anytime it is updated by `fetchComplaints()`, the browser automatically updates the DOM without requiring a full page reload. Also point out how `statusClass(item.status)` dynamically assigns CSS classes to color-code the "Pending" / "Resolved" badges.
