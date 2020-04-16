// author: ashing
// time: 2020/4/16 11:20 上午
// mail: axingfly@gmail.com
// Less is more.

package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
