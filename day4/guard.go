package main

type guards []guard

type guard struct {
	id     int
	shifts shifts
}

func (gs guards) last() *guard {
	return &gs[len(gs)-1]
}
