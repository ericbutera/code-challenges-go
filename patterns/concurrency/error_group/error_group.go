package error_group

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

func ProcessTasks(ctx context.Context, tasks []func(ctx context.Context) error) error {
	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		errors []error
	)
	g, ctx := errgroup.WithContext(ctx)

	for _, task := range tasks {
		wg.Add(1)
		t := task // create a local copy to avoid data race
		g.Go(func() error {
			defer wg.Done()
			err := t(ctx)
			if err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
			}
			return nil
		})
	}

	wg.Wait()

	if len(errors) > 0 {
		return fmt.Errorf("multiple errors: %v", errors)
	}

	return g.Wait()
}
