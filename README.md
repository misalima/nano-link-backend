# Nano Link Backend

A URL shortener service backend built with Go. This service allows users to create shortened URLs, customize them, tag them, and track visits.

## Project Overview

Nano Link is a RESTful API that provides the following features:
- URL shortening with automatically generated or custom short IDs
- User authentication and management
- Tagging system for URLs
- Visit tracking and statistics

## Tech Stack

- Go (Golang)
- PostgreSQL
- Docker
- Echo (Web Framework)
- [Tern](https://github.com/jackc/tern) (Database Migration Tool)

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose
- Tern migration tool

### Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/misalima/nano-link-backend.git
   cd nano-link-backend
   ```

2. **Create environment file**

   Create a `.env` file in the root directory based on the provided `.env.example`:

   ```bash
   cp .env.example .env
   ```

3. **Export environment variables**

   ```bash
   source .env
   ```

   > **Note for Windows users**: If you encounter issues with sourcing the `.env` file, you may need to convert it to Unix format using dos2unix:
   > ```bash
   > dos2unix .env
   > ```
   > You can find more information about dos2unix in the [dos2unix documentation](https://waterlan.home.xs4all.nl/dos2unix.html).

4. **Start the PostgreSQL database**

   ```bash
   docker-compose up -d
   ```

   This will create and start a PostgreSQL container with the configuration specified in your `.env` file.

5. **Run database migrations**

   Install [Tern](https://github.com/jackc/tern) if you haven't already:

   ```bash
   go install github.com/jackc/tern@latest
   ```

   Run the migrations:

   ```bash
   cd config/postgres
   tern migrate
   ```

### Running the Application

Start the application:

```bash
go run src/app/api/main.go
```

The server will start on the host and port specified in your `.env` file (default: localhost:3333).

## API Endpoints

The API provides endpoints for:
- User management
- URL shortening and management
- Tag management
- Visit tracking

For detailed API documentation, refer to the API documentation (not included in this README).

## Development

### Project Structure

- `src/app/api`: API layer (handlers, router, configuration)
- `src/core`: Domain models and business logic
- `src/infra`: Infrastructure layer (database repositories)
- `config`: Configuration files including database migrations

### Database Migrations

Migrations are managed with Tern and stored in `config/postgres/`.

To create a new migration:

```bash
cd config/postgres
tern new migration_name
```

To apply migrations:

```bash
tern migrate
```


