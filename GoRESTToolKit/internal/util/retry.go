package util

import (
	"errors"
	"math"
	"time"
)

func Retry(attempts int, sleep time.Duration, f func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		if err = f(); err == nil {
			return nil
		}
		time.Sleep(time.Duration(math.Pow(2, float64(i))) * sleep)
	}

	return errors.New("all retry attempts failed")
}
