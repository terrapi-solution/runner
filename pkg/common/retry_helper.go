package common

import (
	"fmt"
	"time"
)

type retryHelper struct {
	Retry     int           `json:"retry" description:"How many times to retry the operation"`
	RetryTime time.Duration `json:"retry-time" description:"How long to wait between retries"`
}

type retryableErr struct {
	err error
}

func (e retryableErr) Unwrap() error {
	return e.err
}

func (e retryableErr) Error() string {
	return e.err.Error()
}

func (r *retryHelper) doRetry(handler func(int) error) error {
	err := handler(0)

	for retry := 1; retry <= r.Retry; retry++ {
		if _, ok := err.(retryableErr); !ok {
			return err
		}

		time.Sleep(r.RetryTime)
		fmt.Printf("Attempt %d failed with error: %v. Retrying...\n", retry, err)

		err = handler(retry)
	}

	return err
}
