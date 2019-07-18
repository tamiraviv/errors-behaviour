package errors

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type behaviour int

const (
	Unkown behaviour = iota
	BadRequest
	InternalError
	Temporary
)

type behaviourError struct {
	error
	behaviours map[behaviour]bool
}

func New(message string) *behaviourError {
	return &behaviourError{
		errors.New(message),
		map[behaviour]bool{},
	}
}

func Wrap(err error, message string) *behaviourError {
	var behaviours map[behaviour]bool

	be, ok := err.(*behaviourError)
	if ok {
		behaviours = be.behaviours
	} else {
		behaviours = map[behaviour]bool{}
	}

	return &behaviourError{
		errors.Wrap(err, message),
		behaviours,
	}
}

func Cause(err behaviourError) *behaviourError {
	return nil
}

func (e *behaviourError) Error() string {
	return e.error.Error()
}

func (e *behaviourError) AddBehaviour(b behaviour) *behaviourError {
	e.behaviours[b] = true
	return e
}

func (e *behaviourError) IsTemporary() bool {
	return e.behaviours[Temporary]
}

func (e *behaviourError) IsBadRequest() bool {
	return e.behaviours[BadRequest]
}

func (e *behaviourError) IsInternalError() bool {
	return e.behaviours[InternalError]
}

func (e *behaviourError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", e.error)
			io.WriteString(s, e.error.Error())
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.error.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.error.Error())
	}
}
