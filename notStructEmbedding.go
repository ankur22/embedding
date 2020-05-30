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
	l *logger
	c *checker
}

func main() {
	var lc loggerAndChecker
	in := "Hello World"
	if lc.c.validate(in) {
		lc.l.do(in)
	}
}
