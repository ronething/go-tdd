// author: ashing
// time: 2020/4/16 10:55 上午
// mail: axingfly@gmail.com
// Less is more.

package main

import "fmt"

const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
