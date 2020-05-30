package main

import (
	"log"
)

type logger struct{}

func (l *logger) do(in string) {
	if l.validate(in) {
		log.Println(in)
	}
}

func (l *logger) validate(in string) bool {
	if len(in) == 0 {
		return false
	}
	return true
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
}

func (l *loggerAndChecker) validate(in string) bool {
	return l.logger.validate(in)
}

func main() {
	var lc loggerAndChecker
	in := "Hello World"
	if lc.validate(in) {
		lc.do(in)
	}
}
