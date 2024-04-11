# JP Note

It helps to remember japanese notes.

## Install

```bash
go install github.com/hellflame/jp-note@latest
```

An executable named `jp-note` will be install at `$GOPATH/bin`

## Usage

```
usage: jp-note [--help] [--host HOST] [--port PORT] [--no-browser]

start jp note server

options:
  --help, -h            show this help message
  --host HOST           host address (default: 127.0.0.1)
  --port PORT, -p PORT  listen port (default: 8080)
  --no-browser          don't open browser
```