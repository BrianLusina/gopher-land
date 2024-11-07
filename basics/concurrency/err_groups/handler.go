package errgroups

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type (
	Circle struct{}

	Result struct{}
)

func handler(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	group, ctx := errgroup.WithContext(ctx)

	for i, circle := range circles {
		i := i
		circle := circle

		group.Go(func() error {
			result, err := foo(ctx, circle)
			if err != nil {
				return err
			}
			results[i] = result
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return nil, err
	}

	return results, nil
}

func foo(ctx context.Context, circle Circle) (Result, error) {
	return Result{}, nil
}
