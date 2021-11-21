package a

import (
	"errors"
	aa "github.com/pkg/errors"
)

func A() error {

	return aa.Wrap(errors.New("haaaa"), "aaa123")
}

func B() error {
	err := A()
	if err != nil {

		return aa.Wrap(err, "A failed")
	}
	return nil
}
