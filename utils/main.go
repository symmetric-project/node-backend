package utils

import (
	"github.com/ztrue/tracerr"
)

func Stacktrace(err error) {
	err = tracerr.Wrap(err)
	tracerr.Print(err)
}
