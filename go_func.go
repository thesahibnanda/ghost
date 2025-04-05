package ghost

import "context"

func Go(ctx context.Context, fn func(ctx context.Context)) {
	go func() {
		defer Track(ctx)()
		fn(ctx)
	}()
}
