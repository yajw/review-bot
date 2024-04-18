package logger

import (
	"fmt"
	"os"
)

var Info = func(format string, a ...any) {
	fmt.Fprintf(os.Stdout, format, a)
}

var Error = func(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format, a)
}
