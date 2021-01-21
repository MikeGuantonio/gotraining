# gotraining

Work was done to refresh go following the TDD principals found on https://quii.gitbook.io/learn-go-with-tests/

## Working in container only

### Build container

```bash
docker-compose build
```

### Run tests with container

```bash
docker-compose run golang-train go test <directory or filename>
```

### End Session

```bash
docker-compose down
```