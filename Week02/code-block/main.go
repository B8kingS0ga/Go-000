package main

import (
	"fmt"
	"github.com/pkg/errors"
)

var errMy = errors.New("my")

func test0() error {
	return errors.Wrapf(errMy, "test0 failed")
}
func test1() error {
	return test0()
}

func test2() error {
	return test1()
}

func main() {
	err := test2()

	if errors.Is(err, errMy) {

		fmt.Printf("main: %+v\n", err)
	}

}
