package slice

func Shrink[S ~[]E, E any](s S) S {
	length, capacity := len(s), cap(s)
	ans := calcShrinkCap(length, capacity)
	if ans == 0 {
		return s
	}
	newS := make([]E, 0, ans)
	newS = append(newS, s...)
	return newS
}

func calcShrinkCap(len, cap int) int {
	if cap <= 64 {
		return 0
	}
	if cap >= 2048 && cap >= len<<1 {
		factor := 0.6
		return int(float64(cap) * factor)
	}
	if cap <= 2048 && cap >= len<<2 {
		return cap >> 1
	}
	return 0
}
