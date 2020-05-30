package main

import (
	"log"
)

type doer interface {
	do(string)
}

type validator interface {
	validate(string) bool
}

type doerAndValidator interface {
	doer
	validator
	doSomething(string)
}

type stdoutPrinter struct{}

func (s *stdoutPrinter) do(in string) {
	log.Println(in)
}

func (s *stdoutPrinter) validate(in string) bool {
	if len(in) == 0 {
		return false
	}
	return true
}

func (s *stdoutPrinter) doSomething(in string) {
	log.Println(in)
}

func justDo(d doer, in string)            { d.do(in) }
func justValidate(v validator, in string) { log.Printf("%v\n", v.validate(in)) }
func validateThenDo(dv doerAndValidator, in string) {
	if dv.validate(in) {
		dv.do(in)
	}
}

func main() {
	var p *stdoutPrinter
	justDo(p, "Hello")
	justValidate(p, "Hello")
	validateThenDo(p, "World")
}
