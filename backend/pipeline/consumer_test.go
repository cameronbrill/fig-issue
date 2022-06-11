package pipeline

import (
	"context"
	"testing"
)

func TestConsumer(t *testing.T) {
	input := []string{"FOO", "BAR", "BAX", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}
	ctx, cancel := context.WithCancel(context.Background())
	inChan := Producer(ctx, input)

	errChan := make(chan error)
	defer close(errChan)
	noop := func(s string) error { return nil }
	err := Consumer(ctx, cancel, inChan, noop, errChan)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("canceled context", func(t *testing.T) {
		three := Producer(ctx, input)

		ctx, cancel = context.WithCancel(context.Background())
		cancel()

		err = Consumer(ctx, cancel, three, noop, errChan)
		if err != context.Canceled {
			t.Fatal(err)
		}
	})
}
