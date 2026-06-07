# Member 3: User Add Complaint

## 📌 Module Overview
This module gives logged-in standard users the ability to log new customer complaints into the CRM.

## 🔗 Connections & Dependencies
- **Authentication**: Extracts the `Authorization` header token (provided by Member 2). If the role is 'admin', it returns `403 Forbidden` because admins cannot log complaints.
- **API Route**: Hits `POST http://localhost:8080/api/complaints`.
- **Database**: Runs an `INSERT` statement into the schema defined by Member 1.
- **Triggers**: Upon success, it automatically invokes `this.fetchComplaints()` created by Member 5 to refresh the table UI dynamically.

## 💻 Core Code Explanation

### 1. Backend POST Logic (`backend_add_complaint.go`)
Decodes the JSON payload, checks the user token, enforces role-based access control, and executes a SQL `INSERT`.
```go
if user.Role == "admin" {
    http.Error(w, "Admins may not submit complaints", http.StatusForbidden)
    return
}
c.CustomerID = user.ID // Prevent users from assigning complaints to others

err = db.QueryRow(
    "INSERT INTO complaints (customer_id, title, description, status) VALUES ($1, $2, $3, 'Pending') RETURNING id",
    c.CustomerID, c.Title, c.Description,
).Scan(&c.ID)
```

### 2. Frontend Submission Logic (`frontend_add_complaint.vue`)
Compiles the Vue reactive form data (`this.form`) and sends it over to the REST API with the required authentication headers.
```javascript
async submitComplaint() {
    const res = await fetch(this.apiBase, {
        method: "POST",
        headers: this.getAuthHeaders(), // Function sending the token
        body: JSON.stringify(this.form),
    });
    if (res.ok) {
        this.form = { customer_id: null, title: "", description: "" }; // Clear form
        await this.fetchComplaints(); // Refresh data table
    }
}
```

## 🚀 Viva Presentation Notes
Focus on explaining **Data Validation** and **Authorization Security**. Point out how the backend overrides the requested `CustomerID` with the `user.ID` from the token, preventing malicious users from spoofing complaints on behalf of other customers.
