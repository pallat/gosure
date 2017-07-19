package sure

import (
	"testing"
	"time"
)

type fakeDoOK struct{}

func (f *fakeDoOK) Do() error {
	return nil
}

func TestItsAbsolutelyOK(t *testing.T) {
	ok := Sure(&fakeDoOK{}, 3, 500*time.Millisecond)
	if !ok {
		t.Error("it should be ok")
	}
}
