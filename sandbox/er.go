package sandbox

import "fmt"

type myError struct {
	code int
}

func (e myError) Error() string {
	return fmt.Sprintf("My error %d", e.code)
}

func safeFunc() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)

		}
	}()

	mightPanic()
}

func mightPanic() {
	fmt.Println("About to panic!")
	panic("Something wrong")
}
