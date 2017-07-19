package sure

import (
	"time"
)

type Worker interface {
	Do() error
}

func Sure(w Worker, repeat int64, frequency time.Duration) bool {
	tick := time.Tick(frequency)
	timeout := time.After(time.Duration(repeat*frequency.Nanoseconds()) * time.Nanosecond)
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
		case <-timeout:
			return false
		}
	}
	return false
}
