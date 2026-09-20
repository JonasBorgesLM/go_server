# go_server

**A Go REST API for user management and authentication, built to practice Clean
Architecture with the boundaries actually enforced** — JWT auth over PostgreSQL,
SQL migrations, layered packages, a test suite, CI and SonarQube analysis.

## Why this project

The interesting constraint is the dependency rule, and it is visible in the
imports. `internal/domain` imports nothing from the project. `internal/usecase`
imports the domain and the repository interface, never `database/sql` and never
`net/http`. Only `internal/repository` knows about PostgreSQL; only
`internal/handler` knows about HTTP.

That means the direction of coupling is checkable rather than aspirational: if
`usecase` ever imports `pkg/database`, the architecture broke, and you can see
it in one line of a diff.

## Layout

```
cmd/api/              # entrypoint
internal/
├── domain/           # the user entity — no dependencies
├── dto/              # request/response shapes, with validation
├── formatter/        # dto <-> domain translation, incl. password hashing
├── usecase/          # application logic (user, login)
├── repository/       # PostgreSQL implementation
├── handler/          # HTTP handlers
├── middleware/       # JWT auth middleware
├── server/           # routing and wiring
├── config/           # env loading and validation
└── migrations/       # golang-migrate SQL files
pkg/
├── database/         # connection
├── utils/            # jwt, password hashing, error responses
└── validations/      # email, username, password rules
test/unit/            # unit tests, mirroring internal/
```

`formatter` is the piece worth noticing: DTOs never become domain objects by
assignment. The translation is explicit and one-way, which is where password
hashing happens so a plaintext password never reaches the domain or the
repository.

## Requirements

- Go 1.24+
- PostgreSQL
- Docker (optional, for the database)

## Configuration

Copy `.env.example` to `.env`. Every variable below is **required** — the app
validates them on boot and refuses to start if one is missing or malformed,
rather than failing later on the first request:

| Variable | Purpose |
| --- | --- |
| `DB_HOST`, `DB_PORT`, `DB_USERNAME`, `DB_PASSWORD`, `DB_DATABASE` | PostgreSQL connection (`DB_PORT` must parse as a number) |
| `JWT_SECRET` | Signing key for access tokens |
| `PORT` | HTTP listen port (defaults to `8080`) |

> The `JWT_SECRET` in `.env.example` is a placeholder. Generate your own —
> `openssl rand -base64 32` — and never commit it.

## Running

```bash
docker compose up -d          # PostgreSQL
migrate -path internal/migrations -database "$DATABASE_URL" up
go run ./cmd/api
```

## API

`POST /login` is public. Every other route sits behind `AuthMiddleware` and
requires `Authorization: Bearer <token>`.

| Intended method | Route | Purpose |
| --- | --- | --- |
| `POST` | `/login` | Exchange credentials for a JWT |
| `GET` | `/users` | List users |
| `POST` | `/users/create` | Create a user |
| `PUT` | `/users/update?id={uuid}` | Update a user |
| `DELETE` | `/users/delete?id={uuid}` | Delete a user |
| `GET` | `/users/username?username={name}` | Find a user by username |

> **The method column is intent, not enforcement.** Routes are registered with
> plain `ServeMux` patterns and no handler inspects `r.Method`, so every verb
> reaches every handler — `GET /users/delete?id=...` deletes the user. Fixing
> this is the same change as fixing the route shapes below: Go 1.22 method
> patterns.

Ready-to-run requests live in `http/` (`login.http`, `user.http`, `app.http`).

```bash
curl -X POST http://localhost:8080/login \
  -H 'Content-Type: application/json' \
  -d '{"username":"jonas","password":"..."}'
```

### A note on the route shapes

The verb-in-path style (`/users/create`, `/users/delete`) is not REST — the
resource-oriented form would be `POST /users` and `DELETE /users/{id}`. It is
what the routing here does today, documented as-is rather than as it should be.
Go 1.22's `ServeMux` supports method and wildcard patterns
(`mux.HandleFunc("DELETE /users/{id}", ...)`), which fixes the shape and the
missing method check at once, without adding a router dependency.

## Validation

Input rules live in `pkg/validations` and are applied at the DTO boundary before
anything reaches a use case:

- **Email** — format check
- **Username** — length and character rules
- **Password** — strength rules, hashed with bcrypt in `formatter` before storage

## Tests

```bash
go test ./...
go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
```

Unit tests live under `test/unit/`, mirroring the `internal/` tree.

## Quality

- **CI** — `.github/workflows/ci.yml` runs build, tests and lint on every push
- **Lint** — `golangci-lint run` (config in `.golangci.yml`)
- **Static analysis** — SonarQube, configured in `sonar-project.properties`

## Scope

A study project for practicing architecture and delivery discipline, not a
production service. Notably absent: method enforcement on routes (see above),
refresh tokens, token revocation, rate limiting on `/login`, and structured
logging.
