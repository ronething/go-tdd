// author: ashing
// time: 2020/4/18 15:57 下午
// mail: axingfly@gmail.com
// Less is more.

package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
    field := val.Field(0)
	fn(field.String())
}