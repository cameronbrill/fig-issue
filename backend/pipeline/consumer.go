package pipeline

import (
	"context"

	log "github.com/sirupsen/logrus"
)

func Consumer[T any](ctx context.Context, cancelFunc context.CancelFunc, values <-chan T, operation func(T) error, errors <-chan error) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err, ok := <-errors:
			if ok {
				log.Errorf("error: %v", err)
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
			log.Tracef("Consumed: %v", val)
		}
	}
}
