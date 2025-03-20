package iparser

import (
	"fmt"
	"testing"
)

func TestGoDirectiveTimeoutErr(t *testing.T) {
	var syntax = `
	START("123"){
		NO_CHECK_MISS()

		a -> GO(a, "1") -> a(WAIT("1"))
	}
	`
	p, e := Parse(syntax)
	fmt.Println(p, e)
}
