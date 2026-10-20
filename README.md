# ToyHub - Toy E-commerce Backend

ToyHub is a backend REST API for a toy e-commerce application, built with Go (Golang) and PostgreSQL.

## Features

- Product CRUD operations
- User management
- JWT-based authentication
- Middleware for request handling
- Pagination for product listings
- PostgreSQL database integration with SQLX
- Database migrations

## Tech Stack

- Go (Golang)
- PostgreSQL
- SQLX
- JWT
- Git and GitHub

## Project Structure

```
cmd/          Application entry and server setup
config/       Configuration loading
db/           Database connection
domain/       Core domain models
infra/db/     Database infrastructure
migrations/   SQL migration files
product/      Product logic
user/         User logic
repo/         Repository layer
rest/         REST handlers, routes and middlewares
util/         Helper functions
main.go       Application entry point
```

## API Endpoints

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | /users | Register a new user | No |
| POST | /users/login | Login and get JWT token | No |
| GET | /products | List products (with pagination) | No |
| GET | /products/{id} | Get a product by ID | No |
| POST | /products | Create a product | Yes |
| PUT | /products/{id} | Update a product | Yes |
| DELETE | /products/{id} | Delete a product | Yes |

## Getting Started

1. Clone the repository:
```
   git clone https://github.com/MiftahulJannatSiddiqua/toy-hub-ecommerce-backend.git
   cd toy-hub-ecommerce-backend
```
2. Copy `.env.example` to `.env` and fill in your values.
3. Create a PostgreSQL database.
4. Run the database migrations from the `migrations` folder.
5. Start the application:
```
   go run main.go
```

## Author

Miftahul Jannat Siddiqua
