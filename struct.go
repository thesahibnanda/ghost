package ghost

import (
	"context"
	"sync"
	"time"
)

type Context struct {
	ctx         context.Context
	traceID     string
	startTime   time.Time
	rootSpans   []*Span
	currentSpan *Span
	mu          sync.Mutex
}

type Span struct {
	Name     string
	Start    time.Time
	End      time.Time
	Panick   *Panick
	Children []*Span

	parent *Span
}

type Panick struct {
	Stack string
}
