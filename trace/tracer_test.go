package trace_test

import (
	"bytes"
	"testing"

	"github.com/matt-hoiland/blueprints/trace"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := trace.New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
		t.FailNow()
	}
	tracer.Trace("Hello trace package.")
	if buf.String() != "Hello trace package.\n" {
		t.Errorf("Trace should not write '%s'.", buf.String())
	}
}
