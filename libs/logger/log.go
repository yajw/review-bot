package logger

import (
	"fmt"
	"os"
)

var Info = func(format string, a ...any) {
	fmt.Fprintf(os.Stdout, format+"\n", a)
}

var Error = func(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", a)
}
