# Member 1: Project Setup & Database Initialization

## 📌 Module Overview
This module is the foundational layer of the application. It handles the deployment, environment configuration, database structure, and dependencies. Without this module, the application cannot be containerized or launched.

## 🔗 Connections & Dependencies
- **Database Initialization (`init.sql`)**: Creates the schemas (`complaints` table) that Members 3, 4, and 5 use to perform CRUD operations.
- **Docker Orchestration (`docker-compose.yml`)**: Connects the `frontend`, `backend`, and `database` services together via a virtual Docker network.
- **Dependencies (`go.mod`, `package.json`)**: Installs all required backend and frontend packages, including Vite for the Vue frontend and `pq` (PostgreSQL driver) for the Go backend.

## 💻 Core Code Explanation

### 1. Database Schema (`init.sql`)
This creates the main `complaints` table to store all system complaints.
```sql
CREATE TABLE IF NOT EXISTS complaints (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    status VARCHAR(50) DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 2. Docker Compose Orchestration (`docker-compose.yml`)
Orchestrates three services: the database, the Go backend API, and the Vue frontend application.
```yaml
services:
  db:
    image: postgres:15
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: crm_db
    volumes:
      - ./database/init.sql:/docker-entrypoint-initdb.d/init.sql
  backend:
    build: ./backend
    ports:
      - "8080:8080"
    environment:
      DATABASE_URL: postgres://user:password@db:5432/crm_db?sslmode=disable
  frontend:
    build: ./frontend
    ports:
      - "5173:5173"
```

## 🚀 Viva Presentation Notes
When presenting, explain how `docker-compose up` sets up the entire application. Mention that you defined the initial state of the database so that the team could immediately start developing API endpoints against a structured schema.
