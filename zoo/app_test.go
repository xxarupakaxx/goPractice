package main

import "testing"

func testAppName(t *testing.T) {
	expect := "Zoo Application"
	actual := App()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}

}
