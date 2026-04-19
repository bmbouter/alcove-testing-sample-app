# alcove-testing-sample-app

A sample Go application for testing Alcove's multi-repo SDLC pipelines.

## Dev Container

When a dev container is available, run all commands via the dev container exec
endpoint. Do NOT run go commands directly.

Check if available: `$DEV_CONTAINER_HOST` and `$DEV_TOKEN` env vars are set.

Health check:
```bash
curl -s http://$DEV_CONTAINER_HOST/healthz
```

Run commands:
```bash
curl -s -X POST http://$DEV_CONTAINER_HOST/exec \
  -H "Authorization: Bearer $DEV_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"cmd": "cd /workspace/alcove-testing-sample-app && go test -v ./...", "timeout": 30}'
```

## Testing

Run from the dev container:
```bash
cd /workspace/alcove-testing-sample-app && go test -v ./...
```
