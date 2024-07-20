package sandbox

import "fmt"

type myError struct {
	code int
}

func (e myError) Error() string {
	return fmt.Sprintf("My error %d", e.code)
}

