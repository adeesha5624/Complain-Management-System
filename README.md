# Complain Management App

A simple complaint management system built with:
- Go backend with PostgreSQL
- Vue.js frontend
- Docker Compose to run backend, frontend, and database together

## Project Structure

- `backend/` - Go API server
- `frontend/` - Vue.js single-page application
- `database/` - PostgreSQL initialization script
- `docker-compose.yml` - Docker Compose configuration
- `.env` - environment variables for database and backend connection

## Features

- Admin and user authentication
- Complaint listing and creation
- Admin-only complaint status update and deletion
- Admin UI hides complaint submission form
- PostgreSQL database seeded with demo complaints

## Prerequisites

- Docker Desktop
- Docker Compose

## Setup

1. Copy `.env.example` to `.env` if needed, or create `.env` with values:

```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=crm_db
DATABASE_URL=postgres://postgres:postgres@db:5432/crm_db?sslmode=disable
```

2. Build and start the application:

```bash
docker-compose up --build
```

3. Open the frontend:

- `http://localhost:3000`

## Demo Credentials

- Admin: `admin` / `admin123`
- User: `user` / `user123`

## Notes

- Admin users can view and manage complaints but cannot submit new ones.
- Regular users can submit complaints.
- The backend uses a demo token format and is not production-ready.

## Backend Dev Commands

From `backend/`:

```bash
go build Main.go
./Main
```

## Frontend Dev Commands

From `frontend/`:

```bash
npm install
npm run dev
```
