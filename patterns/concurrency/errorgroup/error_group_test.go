package errorgroup_test

// note: use a real library like https://github.com/sourcegraph/conc

import (
	"context"
	"testing"

	"github.com/ericbutera/code-challenges-go/patterns/concurrency/errorgroup"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessTasks(t *testing.T) {
	t.Parallel()
	t.Run("all tasks succeed", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		tasks := []func(ctx context.Context) error{
			func(_ context.Context) error { return nil },
			func(_ context.Context) error { return nil },
			func(_ context.Context) error { return nil },
		}

		err := errorgroup.ProcessTasks(ctx, tasks)
		assert.NoError(t, err)
	})

	t.Run("one task fails", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		tasks := []func(ctx context.Context) error{
			func(_ context.Context) error { return nil },
			func(_ context.Context) error { return errorgroup.ErrTaskFailed },
			func(_ context.Context) error { return nil },
		}

		err := errorgroup.ProcessTasks(ctx, tasks)
		require.Error(t, err)
	})

	t.Run("all tasks fail", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		tasks := []func(ctx context.Context) error{
			func(_ context.Context) error { return errorgroup.ErrTaskFailed },
			func(_ context.Context) error { return errorgroup.ErrTaskFailed },
			func(_ context.Context) error { return errorgroup.ErrTaskFailed },
		}

		err := errorgroup.ProcessTasks(ctx, tasks)
		require.Error(t, err, "All tasks failed, so an error should be returned")
	})
}
