# User Auth Service

This project provides an authentication service with JWT token-based authentication using Go, PostgreSQL, and Docker. It includes endpoints for user signup, signin, token generation, verification and renewal.

## Table of Contents

- [Requirements](./requirements.md)
- [DB Schema](./schema.md)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Endpoints and CURLs](#api-endpoints)

## Tech Stack

- **Go** - Programming language used for the backend logic.
- **PostgreSQL** - Relational database for storing user and token data.
- **JWT (JSON Web Tokens)** - Used for secure token-based authentication.
- **Docker & Docker Compose** - Containerization and orchestration of services.
- **GORM** - ORM library for interacting with the PostgreSQL database.

## Installation

### Prerequisites

- Docker and Docker Compose
- Go (v1.23+)

### Clone the Repository

```bash
git clone https://github.com/Garg-Shashwat/Go-User-Authentication.git
```

### Installation

```bash
go mod tidy
```

## Configuration

.env.main file included

## Running the Application

```bash
docker-compose up

docker-compose down
```

## API Endpoints

### Ping server

**Endpoint:** `GET: "/ping"`
```bash
curl --location 'http://localhost:9000/ping'
```

### Register User

**Endpoint:** `POST: "/register"`
```bash
curl --location 'http://localhost:9000/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "user@mail.com",
    "password": "Secret"
}'
```

### Login User

**Endpoint:** `POST: "/login"`
```bash
curl --location 'http://localhost:9000/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "user@mail.com",
    "password": "Secret"
}'
```

### Renew Token

**Endpoint:** `GET: "/refresh_token"`
```bash
curl --location 'http://localhost:9000/refresh_token' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI0NzUxOTYsImlhdCI6MTczMjQ2Nzk5Niwic3ViIjoyfQ.l9zXzEt-oEpqOpx5x2Ndxra-cnNA3uj6KW8LDzdkiK8'
```

### Logout (Revoke Token)

**Endpoint:** `GET: "/logout"`
```bash
curl --location 'http://localhost:9000/logout' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzI0NzUxOTYsImlhdCI6MTczMjQ2Nzk5Niwic3ViIjoyfQ.l9zXzEt-oEpqOpx5x2Ndxra-cnNA3uj6KW8LDzdkiK8'
```