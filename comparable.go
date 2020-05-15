package data_structrue

// 比较 Key 的大小
// if x < y return -1
// if x == y return 0
// if x > y return 1
type Comparable func(x, y interface{}) int

func StringComparable(x, y interface{}) int {
	s1 := x.(string)
	s2 := y.(string)

	min := len(s1)
	if len(s2) < len(s1) {
		min = len(s2)
	}

	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}

	if diff == 0 {
		diff = len(s1) - len(s2)
	}

	if diff < 0 {
		return -1
	}

	if diff > 0 {
		return 1
	}
	return 0
}

func IntComparable(x, y interface{}) int {
	i := x.(int)
	j := y.(int)
	if i < j {
		return -1
	} else if i == j {
		return 0
	} else {
		return 1
	}
}