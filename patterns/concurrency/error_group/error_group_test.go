package error_group_test

// note: use a real library like https://github.com/sourcegraph/conc

import (
	"context"
	"errors"
	"testing"

	"github.com/ericbutera/code-challenges-go/patterns/concurrency/error_group"
	"github.com/stretchr/testify/assert"
)

func TestProcessTasks(t *testing.T) {
	t.Run("all tasks succeed", func(t *testing.T) {
		ctx := context.Background()
		tasks := []func(ctx context.Context) error{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return nil },
		}

		err := error_group.ProcessTasks(ctx, tasks)
		assert.NoError(t, err)
	})

	t.Run("one task fails", func(t *testing.T) {
		ctx := context.Background()
		tasks := []func(ctx context.Context) error{
			func(ctx context.Context) error { return nil },
			func(ctx context.Context) error { return errors.New("task failed") },
			func(ctx context.Context) error { return nil },
		}

		err := error_group.ProcessTasks(ctx, tasks)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "task failed")
	})

	t.Run("all tasks fail", func(t *testing.T) {
		ctx := context.Background()
		tasks := []func(ctx context.Context) error{
			func(ctx context.Context) error { return errors.New("task 1 failed") },
			func(ctx context.Context) error { return errors.New("task 2 failed") },
			func(ctx context.Context) error { return errors.New("task 3 failed") },
		}

		err := error_group.ProcessTasks(ctx, tasks)
		assert.Error(t, err, "All tasks failed, so an error should be returned")
		assert.Contains(t, err.Error(), "task 1 failed")
		assert.Contains(t, err.Error(), "task 2 failed")
		assert.Contains(t, err.Error(), "task 3 failed")
	})
}
