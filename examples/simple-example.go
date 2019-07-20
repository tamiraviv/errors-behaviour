package main

import (
	"fmt"

	"errors-behavior/errors"
)

type Temporary interface {
	IsTemporary() bool
}

func main() {
	err := a()
	if err != nil {
		behaviourError, ok := err.(Temporary)
		if ok && behaviourError.IsTemporary() {
			fmt.Printf("Got the following temporary error: %s\n", behaviourError)
		} else {
			fmt.Printf("Got the following unkown error: %s\n", err)
		}
	}
}

func a() error {
	if err := b(); err != nil {
		return errors.Wrap(err, "Error in function a")
	}
	return nil
}

func b() error {
	return errors.New("Error in function b").AddBehaviour(errors.Temporary)
}
