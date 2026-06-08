# Member 6: System Layout and Configuration

## 📌 Module Overview
This module ties all the fragmented pieces together. It serves as the primary entry point for both the Go backend (Server Initialization, Routing, CORS) and the Vue frontend (CSS styling, Shell architecture).

## 🔗 Connections & Dependencies
- **Backend Routing**: Mounts the handlers created by Members 2, 3, 4, and 5 onto specific API URL paths.
- **CORS Management**: Injects headers to allow the Vue frontend (running on Port 5173) to communicate with the Go backend (running on Port 8080).
- **Frontend Container**: Provides the overarching `<div class="app-container">` that hosts the login views from Member 2 and the dashboard views for Members 3, 4, and 5.

## 💻 Core Code Explanation

### 1. Backend Server & CORS Pipeline (`backend_system_config.go`)
Sets up the listener and handles cross-origin requests.
```go
func main() {
    // Mount all member functions
    http.HandleFunc("/api/auth/login", handleLogin)
    http.HandleFunc("/api/auth/user", handleGetUser)
    http.HandleFunc("/api/complaints", handleComplaints)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func corsHeaders(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
```

### 2. Frontend Application Shell (`frontend_system_layout.vue`)
Sets up the glassmorphism layout, navigation header, and CSS grid structure.
```html
<main class="main-content">
    <header class="app-header">
        <div class="logo">
            <h2>CRM - Complain Management Module</h2>
        </div>
        <div class="user-info">
            <span class="role-badge">{{ currentUser.role }}</span>
            <span class="username">{{ currentUser.username }}</span>
        </div>
    </header>
    <div class="dashboard-grid">
        <!-- Components injected here -->
    </div>
</main>
```

## 🚀 Viva Presentation Notes
For your presentation, explain the concept of **CORS (Cross-Origin Resource Sharing)**. The examiner will likely ask why it is needed. Explain that because the Frontend (`localhost:5173`) and Backend (`localhost:8080`) run on different ports, the browser blocks requests for security reasons unless the Backend explicitly permits it via `Access-Control-Allow-Origin` headers, which you implemented. Furthermore, explain how you structured the global CSS variables (`:root`) to maintain a consistent design system.
