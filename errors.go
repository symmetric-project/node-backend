package main

import (
	"github.com/ztrue/tracerr"
)

func HandleError(err error) {
	err = tracerr.Wrap(err)
	tracerr.Print(err)
}
