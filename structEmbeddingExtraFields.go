package main

import (
	"log"
)

type logger struct {
	s  string
	s2 string
}

func (l *logger) do(in string) {
	log.Println(in)
}

type checker struct {
	s  string
	s2 string
}

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
	lc := loggerAndChecker{logger: &logger{s2: "Hi"}, checker: &checker{s2: "Bye"}, s: "Hello World"}
	in := lc.s
	if lc.validate(in) {
		lc.do(in)
	}

	// Comment out to see compilation error
	// lc.do(lc.s2)
	lc.do(lc.logger.s2)
	lc.do(lc.checker.s2)
}
