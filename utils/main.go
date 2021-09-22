package utils

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"time"

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

func NewOctid() string {
	// Bytes worth 8 characters
	bytes := make([]byte, 4)

	// Randomize bytes
	rand.Read(bytes)

	// Encode the randomized bytes as a string
	return hex.EncodeToString(bytes)
}

func CurrentTimestamp() int {
	return int(time.Now().Unix())
}
