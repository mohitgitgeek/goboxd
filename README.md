# goboxd

Go HTTP service for running untrusted code in containerized sandboxes.

## Build and run

```sh
git submodule update --init --recursive
make build
make run
```

Service health endpoint: `http://localhost:8080/healthz`.

## Development commands

All commands run in Docker:

- `make test`
- `make integration`
- `make lint`
- `make fmt`
- `make save`
- `make load`

## Documentation

- [API](docs/api.md)
- [Languages](docs/languages.md)
- [Security](docs/security.md)
- [Benchmarks](docs/benchmarks.md)
- [Architecture](docs/architecture.md)
