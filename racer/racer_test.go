// author: ashing
// time: 2020/4/18 1:20 上午
// mail: axingfly@gmail.com
// Less is more.

package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "https://blog.ronething.cn"
	fastURL := "https://www.baidu.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
