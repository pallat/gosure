package sure

import (
	"time"
)

type Worker interface {
	Do() error
}

func Sure(w Worker, timeout, frequency time.Duration) bool {
	tick := time.Tick(frequency)
	to := time.After(timeout)
	if err := w.Do(); err == nil {
		return true
	}

	for range tick {
		cherr := make(chan error)

		go func(cherr chan error) {
			err := w.Do()
			if err != nil {
				cherr <- err
				return
			}
			cherr <- nil
		}(cherr)

		select {
		case err := <-cherr:
			if err == nil {
				return true
			}
		case <-to:
			return false
		}
	}
	return false
}
