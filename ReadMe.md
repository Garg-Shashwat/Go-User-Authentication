# User Auth Service

This project provides an authentication service with JWT token-based authentication using Go, PostgreSQL, and Docker. It includes endpoints for user signup, signin, token generation, verification and renewal.

## Table of Contents

- [Requirements](./requirements.md)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)

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