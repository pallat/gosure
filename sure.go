package sure

import (
	"time"
)

type Worker interface {
	Do() error
}

func Sure(w Worker, repeat int, frequency time.Duration) bool {
	return true
}
