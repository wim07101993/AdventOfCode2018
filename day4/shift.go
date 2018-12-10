package main

import "time"

type shifts []shift

type shift struct {
	start        time.Time
	end          time.Time
	sleepPeriods sleepPeriods
}

type sleepPeriods []sleepPeriod

type sleepPeriod struct {
	start time.Time
	end   time.Time
}

func (ss shifts) last() *shift {
	return &ss[len(ss)-1]
}

func (ss sleepPeriods) last() *sleepPeriod {
	return &ss[len(ss)-1]
}
