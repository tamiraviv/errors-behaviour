package main

import (
	"fmt"

	"errors-behavior/errors"
)

type Temporary interface {
	Error() string
	IsTemporary() bool
}

func main() {
	fmt.Println("Starting main")
	err := a()
	if err != nil {
		behaviourError, ok := err.(Temporary)
		if ok {
			fmt.Printf("Got the following error: %s, with temporary behaviour\n", behaviourError)
		} else {
			fmt.Printf("Got the following error: %s, with unkown behaviour\n", err)
		}
	}

	fmt.Println("Finished main")
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
