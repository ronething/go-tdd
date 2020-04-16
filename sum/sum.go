// author: ashing
// time: 2020/4/16 3:53 下午
// mail: axingfly@gmail.com
// Less is more.

package sum

// calc sum
func Sum(num [5]int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}
