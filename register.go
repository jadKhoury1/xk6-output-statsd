// Package statsd registers the extension for output
package statsd

import (
	"github.com/jadKhoury1/xk6-output-statsd/pkg/statsd"
	"go.k6.io/k6/output"
)

func init() {
	output.RegisterExtension("output-statsd", statsd.New)
}
