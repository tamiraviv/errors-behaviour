package main

import (
	"fmt"
	"golang.org/x/xerrors"
)

var tempError = xerrors.New("Temporary Error")

func main() {
	err := c()
	if err != nil {
		if xerrors.Is(err,tempError) {
			fmt.Printf("Got the following temporary error: %s\n", err)
		} else {
			fmt.Printf("Got the following unkown error: %s\n", err)
		}
	}
}

func c() error {
	if err := d(); err != nil {
		return xerrors.Errorf("Error in funcion C: %w", err)
	}
	return nil
}

func d() error {
	return xerrors.Errorf("Error in funcion D: %w", tempError)
}
