package ghost

import "runtime/debug"

func getStack() []byte {
	return debug.Stack()
}
