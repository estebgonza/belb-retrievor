package retrievor

import (
	"strconv"
)

func winIndexString(a []string) int {
	s1, _ := strconv.Atoi(a[0])
	s2, _ := strconv.Atoi(a[1])
	if s1 > s2 {
		return 1
	} else if s2 > s1 {
		return 2
	}
	return 0
}
