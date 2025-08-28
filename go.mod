module github.com/qawolfsocket/qawolf-socket-go-module

go 1.22

require (
	github.com/google/uuid v1.6.0         // UUID generation
	github.com/sirupsen/logrus v1.9.3     // Structured logging
	github.com/stretchr/testify v1.9.0    // Testing helpers (assert, require)
	github.com/pkg/errors v0.9.1          // Error wrapping with stack traces
	golang.org/x/net v0.25.0              // Extra networking utilities
	golang.org/x/sync v0.6.0              // Concurrency primitives (errgroup, etc.)
)
