package helper

import (
	"errors"
	"fmt"
	"time"
)

type RetryHelper struct {
	Retry     int           `json:"retry" description:"How many times to retry the operation"`
	RetryTime time.Duration `json:"retry-time" description:"How long to wait between retries"`
}

type RetryableErr struct {
	Err error
}

func (e RetryableErr) Unwrap() error {
	return e.Err
}

func (e RetryableErr) Error() string {
	return e.Err.Error()
}

func (r *RetryHelper) DoRetry(handler func(int) error) error {
	err := handler(0)

	for retry := 1; retry <= r.Retry; retry++ {
		var retryableErr RetryableErr
		if !errors.As(err, &retryableErr) {
			return err
		}

		time.Sleep(r.RetryTime)
		fmt.Printf("Attempt %d failed with error: %v. Retrying...\n", retry, err)

		err = handler(retry)
	}

	return err
}
