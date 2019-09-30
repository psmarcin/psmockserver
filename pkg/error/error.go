package error

import "fmt"

func Wrap(err error, msg string) error {
	if err != nil {
		return fmt.Errorf("%s %w", msg, err)
	}
	return nil
}
