# X5 Hackathon Backend

## Requirements

- Go 1.14.4
- MongoDB

## Building

```bash
go build ./cmd/backend
```

## Running

1. Firstly, you need to generate config:
    ```bash
    ./backend -genconfig > backend.toml
    ```

2. Then just edit it, and start executable with specified config:
    ```bash
    ./backend -config ./backend.toml
    ```