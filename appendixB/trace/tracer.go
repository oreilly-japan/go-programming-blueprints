package trace

import (
	"fmt"
	"io"
)

// Tracerはコード内での出来事を記録します。
type Tracer struct {
	out io.Writer
}

func (t *Tracer) Trace(a ...interface{}) {
	if t == nil || t.out == nil {
		return
	}
	fmt.Fprintln(t.out, a...)
}

func New(w io.Writer) *Tracer {
	return &Tracer{out: w}
}
