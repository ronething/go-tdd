// author: ashing
// time: 2020/4/18 1:21 上午
// mail: axingfly@gmail.com
// Less is more.

package racer

import (
	"fmt"
	"net/http"
	"time"
)

func httpGetTime(url string) time.Duration {
	s := time.Now()
	resp, _ := http.Get(url)
	d := time.Since(s)
	if resp.StatusCode == 200 {
		fmt.Println(url, " status code is 200")
	}

	return d
}

//func Racer(a, b string) (winner string) {
//	aDuration := httpGetTime(a)
//
//	bDuration := httpGetTime(b)
//
//	if aDuration < bDuration {
//		return a
//	}
//
//	return b
//}

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
