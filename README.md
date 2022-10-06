# Management System Logs Service

This is the logs service for my management system.

# Used Technologies

- Golang
- GoFiber
- Postgresql

# Quick Start

Started Manually

```console
$ export DB_USER="postgres"
$ export DB_PASSWORD="password"
$ export DB_HOST="127.0.0.1"
$ export DB_NAME="logs"
$ export PORT="3004"
$ export ACCOUNTS_SERVICE="http://127.0.0.1:3000"
$ export WORKSPACES_SERVICE="http://127.0.0.1:3001"

$ go run ./src/main.go
```

Using Docker-compose

```
$ docker-compose up --build
```

# Documentation

You can find the documentation for each route in the `docs` directory.
