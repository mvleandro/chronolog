package Level

// LogLevel represents the severity level of a log entry.
// It is a string-based enumerated type used to ensure consistency in log level values
// across the application. Use the predefined constants (e.g., Info, Error)
// instead of raw strings to avoid typos and improve code clarity.
//
// Supported log levels:
//   - Trace: for fine-grained, low-level debugging information.
//   - Debug: for detailed debugging information that is more verbose than info logs.
//   - Info:  for standard informational messages about application behavior.
//   - Warn:  for potentially problematic situations that require attention.
//   - Error: for serious issues indicating failures in execution.
type LogLevel string

const (
	// Trace represents verbose diagnostic logs for tracing execution.
	// These logs are typically used in development or for deep debugging scenarios.
	Trace LogLevel = "trace"

	// Debug is used for detailed debugging information that is more verbose than info logs.
	// It is useful for diagnosing issues during development or troubleshooting.
	Debug LogLevel = "debug"

	// Info indicates normal operational messages that require no action.
	// Use this level to log events like successful startups or expected workflow steps.
	Info LogLevel = "info"

	// Warn signals a condition that is unexpected, but the application can continue running.
	// Use this level for degraded performance, retries, or recoverable issues.
	Warn LogLevel = "warn"

	// Error indicates a serious failure that prevents the application or request from continuing as expected.
	// These logs typically require investigation or alerting.
	Error LogLevel = "error"
)

var LogLevelPriority = map[LogLevel]int{
	"trace": 0,
	"debug": 1,
	"info":  2,
	"warn":  3,
	"error": 4,
}
