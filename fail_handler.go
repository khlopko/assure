package assure

import "fmt"

type FailHandler func(msg string)

// / Default fail handler that will panic with a message.
func PanicHandler() FailHandler {
	return func(msg string) {
		panic(fmt.Sprintf("Assetion failed: %s", msg))
	}
}
