package math

import "time"

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func AbsSlow(a int) int {
	time.Sleep(time.Nanosecond)
	if a > 0 {
		return a
	}
	return -a
}
