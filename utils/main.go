package utils

import (
	"os"

	"github.com/ztrue/tracerr"
)

func Stacktrace(err error) {
	err = tracerr.Wrap(err)
	tracerr.Print(err)
}

func ExitWithStacktrace(err error) {
	err = tracerr.Wrap(err)
	tracerr.Print(err)
	os.Exit(1)
}
