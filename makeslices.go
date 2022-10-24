package makeslices

func MakeSlicesNoCap(n int) {
	s := make([][10]int, 0)
	for i := 0; i < n; i++ {
		s = append(s, [10]int{})
	}
	_ = s
}

func MakeSlicesCap(n int) {
	s := make([][10]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, [10]int{})
	}
	_ = s
}
