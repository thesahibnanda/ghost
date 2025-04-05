package ghost

import (
	"context"
	"runtime"
	"strings"
	"time"
)

func (c *Context) startSpan(name string) *Span {
	span := &Span{
		Name:  name,
		Start: time.Now(),
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.currentSpan != nil {
		c.currentSpan.Children = append(c.currentSpan.Children, span)
		span.parent = c.currentSpan
	} else {
		c.rootSpans = append(c.rootSpans, span)
	}

	c.currentSpan = span
	return span
}

func (c *Context) endSpan(span *Span) {
	c.mu.Lock()
	defer c.mu.Unlock()

	span.End = time.Now()

	if span.parent != nil {
		c.currentSpan = span.parent
	} else {
		c.currentSpan = nil
	}
}

func Track(ctx context.Context, name ...string) func() {

	var funcName string
	if len(name) == 0 {
		funcName = getCallerName()
	} else {
		funcName = name[0]
	}

	gc, ok := From(ctx)
	if !ok {
		return func() {
			// No Op
		}
	}

	span := gc.startSpan(funcName)

	return func() {
		if r := recover(); r != nil {
			span.Panick = &Panick{
				Stack: string(getStack()),
			}
		}
		gc.endSpan(span)
	}
}

func getCallerName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "anonymous"
	}
	parts := strings.Split(fn.Name(), "/")
	return parts[len(parts)-1]
}
