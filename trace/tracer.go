package trace

import (
	"fmt"
	"io"
)

// Tracer is the interfafce that describes an object capable of traicing events throughout code.
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func (t *tracer) Trace(a ...interface{}) {
	_, _ = t.out.Write([]byte(fmt.Sprint(a...)))
	_, _ = t.out.Write([]byte("\n"))
}

type nilTracer struct{}

func Off() Tracer {
	return &nilTracer{}
}

func (_ *nilTracer) Trace(a ...interface{}) {}
