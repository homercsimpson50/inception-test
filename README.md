# inception-test

A test project for containerized Gas Town. The containerized mayor and its polecats will build a simple Go HTTP server here.

## Expected Result

After the containerized GT works on this:
- `main.go` — HTTP server with `GET /health` returning `{"status":"ok"}`
- `main_test.go` — Tests for the health endpoint
- `go.mod` — Go module definition
