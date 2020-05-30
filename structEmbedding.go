package main

import (
	"log"
)

type logger struct{}

func (l *logger) do(in string) {
	log.Println(in)
}

type checker struct{}

func (c *checker) validate(in string) bool {
	if len(in) == 0 {
		return false
	}
	return true
}

type loggerAndChecker struct {
	*logger
	*checker
	s string
}

func main() {
	lc := loggerAndChecker{s: "Hello World"}
	in := lc.s
	if lc.validate(in) {
		lc.do(in)
	}
}
