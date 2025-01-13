package test

import "testing"

func AssertEq(res any, exp any, t *testing.T) {
	if res != exp {
		t.Errorf(
			"Equality assertion failed.\n"+
				"\tRecieved: %v\n"+
				"\tExpected: %v",
			exp,
			res,
		)
	}
}
