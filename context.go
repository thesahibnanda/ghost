package ghost

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func Wrap(ctx context.Context, traceIdList ...string) context.Context {
	traceId := uuid.New().String()
	if len(traceIdList) > 0 {
		traceId = traceIdList[0]
	}

	gCtx := &Context{
		ctx:       ctx,
		traceID:   traceId,
		startTime: time.Now(),
		rootSpans: []*Span{},
	}

	return context.WithValue(ctx, ctxKey, gCtx)
}
func From(ctx context.Context) (*Context, bool) {
	val := ctx.Value(ctxKey)
	if gc, ok := val.(*Context); ok {
		return gc, true
	}
	return nil, false
}
