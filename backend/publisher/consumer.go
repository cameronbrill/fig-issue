package publisher

import (
	"context"
	"log"
)

func Consumer[T any](ctx context.Context, cancelFunc context.CancelFunc, values <-chan T, operation func(T) error, errors <-chan error) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err, ok := <-errors:
			if ok {
				log.Printf("error: %v", err)
				if err.Error() == "comment is not issue" {
					continue
				}
				cancelFunc()
				return err
			}
		case val, ok := <-values:
			if !ok {
				return nil
			}
			err := operation(val)
			if err != nil {
				return err
			}
			log.Printf("Consumed: %v", val)
		}
	}
}
