# fizzbuzz-server

REST API implementing a generalized FizzBuzz, built in Go.

## Overview

Given `int1`, `int2`, `limit`, `str1`, `str2`, the API returns the sequence from
1 to `limit`, replacing multiples of `int1` by `str1`, multiples of `int2` by
`str2`, and multiples of both by `str1str2`.

A bonus `/stats` endpoint tracks the most frequently requested parameter combination.

## Requirements

- Go 1.27.1
- Docker (optional)

## Endpoints

An API description is available in [`openapi.yaml`](./openapi.yaml). You can view it interactively by pasting
its content into https://editor.swagger.io.

### `GET /fizzbuzz`

| Param | Type   | Required | Description              |
|-------|--------|----------|--------------------------|
| int1  | int    | yes      | must be > 0              |
| int2  | int    | yes      | must be > 0              |
| limit | int    | yes      | must be > 0 and <= 10000 |
| str1  | string | yes      | non-empty                |
| str2  | string | yes      | non-empty                |

**Example**

```bash
curl "http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=15&str1=fizz&str2=buzz"
```

```json
["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"]
```

**Edge case:** if `int1 == int2`, every multiple is a multiple of both, so it is replaced by `str1str2`.

**Errors**

Invalid input returns `400 Bad Request`:

```bash
curl "http://localhost:8080/fizzbuzz?int1=abc&int2=5&limit=15&str1=fizz&str2=buzz"
```

```json
{"error": "int1 must be an integer"}
```

```bash
curl "http://localhost:8080/fizzbuzz?int1=3&int2=5&limit=20000&str1=fizz&str2=buzz"
```

```json
{"error": "limit must be between 1 and 10000"}
```

### `GET /stats`

Returns the most frequently requested combination of parameters. Always responds with `200 OK`.

```bash
curl "http://localhost:8080/stats"
```

```json
{"request": {"int1":3,"int2":5,"limit":15,"str1":"fizz","str2":"buzz"}, "hits": 2}
```

If no request has been recorded yet:

```json
{"message": "no requests recorded yet"}
```

### `GET /health`

Basic liveness check. Responds with `200 OK` and `{"status": "ok"}`.

## Running locally

```bash
go run ./cmd/server
```

The server listens on port `8080` by default. Override it with the `PORT`
environment variable:

```bash
PORT=9090 go run ./cmd/server
```

## Running with Docker

```bash
docker build -t fizzbuzz-server .
docker run -p 8080:8080 fizzbuzz-server
```

## Testing

```bash
go test ./... -v -cover
```

On Linux/macOS, or WSL, the race detector can also be used:

```bash
go test ./... -race
```

## Project structure

```
cmd/server/         entry point: server startup
internal/fizzbuzz/  sequence generation (logic)
internal/stats/     request counter
internal/api/       HTTP handlers, routing, request validation
```

`internal/fizzbuzz` and `internal/stats` have no HTTP dependency and are tested in isolation.
`internal/api` translates HTTP requests into calls to these packages.

## Design decisions

- **`limit` is capped at 10000** to prevent large memory allocations from a single request.
- **`stats.Request` is used directly as a map key.** All its fields are comparable, so no serialization is needed to identify a distinct request.
- **`stats.Store` is protected by a `sync.Mutex`.** An early version used an unprotected map. A dedicated concurrency test (100 goroutines x 100 calls) immediately triggered `fatal error: concurrent map writes`, which is not just an inaccurate counter but a process crash. The mutex was added to fix this, and the same test now passes reliably across repeated runs.
- **Only valid `/fizzbuzz` requests are recorded** in the stats store, so the "most frequent request" reflects real API usage rather than malformed input. The trade-off is that invalid traffic is not visible in `/stats`.
- **The stats store is in-memory only.** It is not persisted and does not scale across multiple server instances. In a real production setup, this would be backed by Redis or a database with a bounded key space.