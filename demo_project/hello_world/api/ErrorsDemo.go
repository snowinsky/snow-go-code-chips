package api

import (
	"github.com/pkg/errors"
	"log"
)

// go get github.com/pkg/errors
func ErrorsDemo() {
	fakeErr := errors.New("fake exception")
	log.Println(errors.WithStack(fakeErr))
	log.Println(errors.Cause(fakeErr))
}
