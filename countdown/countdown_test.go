// author: ashing
// time: 2020/4/17 10:39 上午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
//	buffer := &bytes.Buffer{}
//
//	Countdown(buffer)
//
//	got := buffer.String()
//	want := `3
//2
//1
//Go!`
//
//	if got != want {
//		t.Errorf("got '%s' want '%s'", got, want)
//	}


	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
