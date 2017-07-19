package sure

import (
	"errors"
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

type fakeDoError struct {
	times int
}

func (f *fakeDoError) Do() error {
	f.times++
	return errors.New("fail")
}

func TestItsAbsolutelyError(t *testing.T) {
	ok := Sure(&fakeDoError{}, 3, 500*time.Millisecond)
	if ok {
		t.Error("it should be ok")
	}
}

func TestItsAbsolutelyErrorShouldEnsureByDoItThreeTimes(t *testing.T) {
	w := fakeDoError{}
	ok := Sure(&w, 3, 500*time.Millisecond)
	if ok {
		t.Error("it should be ok")
	}
	if w.times != 3 {
		t.Error("it should repeat 3 times but got", w.times)
	}
}
