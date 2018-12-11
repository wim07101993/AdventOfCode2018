package main

import "time"

type guards []guard

type guard struct {
	id     int
	shifts shifts
}

func (gs guards) last() *guard {
	return &gs[len(gs)-1]
}

func (gs guards) findSleepyGuard() *guard {
	var longestSleep time.Duration
	var sleepyGuard *guard
	for _, g := range gs {
		tta := g.totalTimeAsleep()
		if tta > longestSleep {
			longestSleep = tta
			sleepyGuard = &g
		}
	}

	return sleepyGuard
}

func (g *guard) totalTimeAsleep() time.Duration {
	var totalTime time.Duration
	for _, shift := range g.shifts {
		for _, sleep := range shift.sleepPeriods {
			totalTime += sleep.end.Sub(sleep.start)
		}
	}

	return totalTime
}

func (g *guard) getSleepyMinute() time.Time {
	return time.Time{}
}
