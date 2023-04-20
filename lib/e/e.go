package e

import (
	"errors"
	"fmt"
)

var (
	ErrUnknowEventType = errors.New("unknow event type")
	ErrUnknowMetaType = errors.New("unknow meta type")
	ErrNoSavedPages = errors.New("no saved pages")
)

func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func WrapIsErr(msg string, err error) error {
	if err == nil {
		return nil
	}

	return Wrap(msg, err)
}
