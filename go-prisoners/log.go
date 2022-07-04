package main

import (
	"fmt"
	"time"
)

func log(s string, args ...interface{}) {
	datestamp := time.Now().UTC().Format(time.RFC3339)
	fmt.Printf(datestamp+" "+s+"\n", args...)
}
