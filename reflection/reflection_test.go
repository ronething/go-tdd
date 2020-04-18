// author: ashing
// time: 2020/4/18 15:57 下午
// mail: axingfly@gmail.com
// Less is more.

package reflection

import "testing"

func TestWalk(t *testing.T) {

    expected := "Chris"
    var got []string

    x := struct {
        Name string
    }{expected}

    walk(x, func(input string) {
        got = append(got, input)
    })

	if got[0] != expected {
		t.Errorf("got '%s', want '%s'", got[0], expected)
	}
}