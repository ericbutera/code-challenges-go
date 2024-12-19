package errorgroup

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var ErrTaskFailed = errors.New("task failed")

func ProcessTasks(ctx context.Context, tasks []func(ctx context.Context) error) error {
	var (
		waitGroup sync.WaitGroup
		mutex     sync.Mutex
		errors    []error
	)
	group, ctx := errgroup.WithContext(ctx)

	for _, task := range tasks {
		waitGroup.Add(1)
		t := task // create a local copy to avoid data race
		group.Go(func() error {
			defer waitGroup.Done()
			err := t(ctx)
			if err != nil {
				mutex.Lock()
				errors = append(errors, err)
				mutex.Unlock()
			}
			return nil
		})
	}

	waitGroup.Wait()

	if len(errors) > 0 {
		return fmt.Errorf("%w: %v", ErrTaskFailed, errors)
	}

	if err := group.Wait(); err != nil {
		return fmt.Errorf("%w: %v", ErrTaskFailed, errors)
	}
	return nil
}
