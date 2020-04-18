// author: ashing
// time: 2020/4/18 15:57 下午
// mail: axingfly@gmail.com
// Less is more.

package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			// 只有 string 类型才需要进行 fn 调用
			fn(field.String())
		}
	}
}
