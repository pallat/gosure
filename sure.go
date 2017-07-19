package sure

import (
	"time"
)

type Worker interface {
	Do() error
}

func Sure(w Worker, repeat int, frequency time.Duration) bool {
	if err := w.Do(); err != nil {
		return false
	}
	return true
}
