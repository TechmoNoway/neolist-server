# neolist-server

neolist-server is a small Go REST API that provides user and task management functionality.

## Features

- REST API for users and tasks
- Services, repositories, DTOs and handlers organized under `internal/`

## Prerequisites

- Go 1.24.4+ installed
- A SQL database MySQL configured for the application (see your environment configuration)

## Quickstart

Run the server directly with Go:

```bash
go run ./cmd/server
```

Build and run binary:

```bash
go build -o bin/neolist ./cmd/server
./bin/neolist
```

If a `Makefile` is available, you may also have project-specific targets such as `make run`.

## Configuration

The server reads configuration from environment variables and/or the config package. Common variables to set:

- `PORT` — server listening port (default: `8080`)
- `DATABASE_URL` — database connection string

Adjust values to match your local or production environment.

## Tests

Run unit tests:

```bash
go test ./...
```

## API (examples)

The project exposes typical user and task endpoints. Example endpoints:

- `POST /users` — create a user
- `GET /users` — list users
- `POST /tasks` — create a task
- `GET /tasks` — list tasks

Refer to the handlers in `internal/handlers` for exact request/response shapes.

## Contributing

- Create a topic branch for your change (e.g., `feat/add-login`)
- Add or update tests for new behavior
- Open a pull request against `main` with a clear description

## License

No license file is included in the repository. Add a `LICENSE` file if you wish to declare terms for this project.

---

Feel free to ask me to expand any section (detailed env vars, examples, or a Postgres docker-compose). 
