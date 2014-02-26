package errs

import (
	"bytes"
	"fmt"
	"runtime"
)

type trace struct {
	file string
	line int
}

type Err struct {
	err    error
	traces []trace
}

var _ error = new(Err)

func New(f string, args ...interface{}) *Err {
	orig := fmt.Errorf(f, args...)
	err := &Err{
		err:    orig,
		traces: make([]trace, 0),
	}
	err.trace(1)
	return err
}

func Wrap(e error) *Err {
	es, ok := e.(*Err)
	if ok {
		es.trace(1)
		return es
	}

	err := &Err{
		err:    e,
		traces: make([]trace, 0),
	}
	err.trace(1)
	return err
}

func (e *Err) Error() string {
	buf := bytes.NewBufferString(e.err.Error())
	for _, trace := range e.traces {
		fmt.Fprintf(buf, "\n    %s:%d", trace.file, trace.line)
	}
	return buf.String()
}

func (e *Err) trace(skip int) {
	_, file, line, _ := runtime.Caller(skip + 1)
	t := trace{file, line}
	e.traces = append(e.traces, t)
}
