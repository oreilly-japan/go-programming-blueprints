package trace

import "io"

// Tracerはコード内での出来事を記録できるオブジェクトを表すインタフェースです。
type Tracer interface {
	Trace(...interface{})
}

func New(w io.Writer) {}
