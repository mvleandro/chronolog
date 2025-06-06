# Chronolog

**Chronolog** is a lightweight and extensible structured logging library for Go, designed to provide developers with a rich and uniform logging experience across distributed applications.

It supports multiple output formats (JSON and pretty), structured log levels, context propagation for traceability, and common application patterns like operations, events, traces, and errors.

---

## ✨ Features

- ✅ Structured logs with semantic fields
- 🧵 Context-aware logging (trace/span/parent IDs)
- 📦 Extensible log types: `Trace`, `Operation`, `Message`, `Lambda`, etc.
- 📃 Output formats: JSON (machine-friendly) and Pretty (human-friendly)
- 🎚️ Minimum log level filtering
- 🔧 Simple configuration

---

## 🚀 Getting Started

### Installation

```bash
go get github.com/Astronotify/chronolog
```

### Basic Setup

```go
package main

import (
  "context"
  "github.com/Astronotify/chronolog"
)

func main() {
  chronolog.Setup(chronolog.Config{
    Format: chronolog.FormatPretty,
    MinimumLogLevel:  chronolog.Info,
  })

  ctx := context.Background()
  chronolog.Info(ctx, "Service started successfully")
}
```

Call `chronolog.Setup` during your application's initialization phase before
invoking any logging functions. If it isn't called, Chronolog will fall back to
a basic text logger.

---

## 🧱 Log Entry Types

Chronolog provides multiple structured log types:

### 🔹 LogEntry

Base log structure, contains:

- `Timestamp`
- `Level`
- `Message`
- `TraceID`, `SpanID`, `ParentSpanID`
- `Version`, `CommitHash`, `BuildTime`
- `AdditionalData`

### 🔹 ErrorLogEntry

Used with `chronolog.Error(...)`:
```go
chronolog.Error(ctx, fmt.Errorf("database unavailable"))
```

### 🔹 OperationRequest / OperationResponse

```go
req := entries.NewOperationRequestLogEntry(ctx, "create-profile", "/profile", "123", "POST")
res := entries.NewOperationResponseLogEntry(req, 200)
```

### 🔹 MessageReceived / Acknowledged / Rejected

```go
recv := entries.NewMessageReceivedLogEntry(ctx, "msg-id", "topic", "consumer-a")
ack := entries.NewMessageAcknowledgedLogEntry(ctx, "msg-id", "topic", "consumer-a")
rej := entries.NewMessageRejectedLogEntry(ctx, "msg-id", "topic", "consumer-a", "invalid payload")
```

### 🔹 TraceBegin / TraceEnd

```go
begin := entries.NewTraceBeginLogEntry(ctx, "cache-warmup")
end := entries.NewTraceEndLogEntryFromBegin(begin)
```

---

## 📦 Output Formats

Chronolog supports:

### JSON Format (default)
```json
{
  "timestamp": "2025-06-01T14:22:10Z",
  "level": "info",
  "event_type": "LogEntry",
  "message": "Service initialized"
}
```

### Pretty Format (human-readable)
```
[2025-06-01T14:22:10Z] [INFO] Service initialized trace_id=abc123 span_id=def456 ...
```

To configure format:

```go
chronolog.Setup(chronolog.Config{
  Format: chronolog.FormatPretty,
})
```

---

## 🛠 Configuration

```go
chronolog.Setup(chronolog.Config{
  Writer: os.Stdout,
  Format: chronolog.FormatJSON, // or FormatPretty
  MinimumLogLevel:  chronolog.Info,
})
```

### Minimum Log Level

Logs below the configured level will be discarded.

| Level Order | Name  |
|-------------|-------|
| 1           | trace |
| 2           | debug |
| 3           | info  |
| 4           | warn  |
| 5           | error |

---

## 🌐 Context Enrichment

Use the helpers in the `ctx` package to add trace information to your logs:

```go
import (
    chronologctx "github.com/Astronotify/chronolog/ctx"
)

ctx := context.Background()
ctx = chronologctx.WithTraceID(ctx, "trace-12345")
chronolog.Info(ctx, "processing request")
```

---

## 📂 Project Structure

```
chronolog/
├── entries/         # Log entry types
├── ctx/             # Public context helpers
├── level/           # Log level definitions
├── internal/        # Utility and handler logic (internal use only)
├── chronolog.go     # Main API
└── README.md
```

---

## 📚 Examples

Several runnable examples live in the `examples/` directory:

- `main.go` – comprehensive demo covering traces, operations, messages and more.
- `httpserver/` – basic HTTP handler emitting `OperationRequestLogEntry` and
  `OperationResponseLogEntry` for each request.
- `message_consumer/` – illustrates a message lifecycle with `Received`,
  `Acknowledged` and `Rejected` events.
- `lambda/` – shows how to pair `LambdaBeginLogEntry` and `LambdaEndLogEntry`.

Run any example with `go run ./examples/<name>`.

---

## 📖 License

MIT License
