package utils

func EscapeToTime(ok bool, time int) int {
	if ok {
		return time
	} else {
		return -1
	}
}
