package main

import (
	"fmt"
	"time"
)

type Todo struct {
	id           int64
	done         bool
	summary      string
	description  string
	deadline     time.Time
	scheduledFor time.Time
}

func main() {
	fmt.Printf("Hello World")
}
