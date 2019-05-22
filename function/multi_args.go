package funcstudy

func sum(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}
	return s
}

func sum2(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
