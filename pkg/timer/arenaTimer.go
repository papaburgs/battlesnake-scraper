package timer

import "time"

func ScrapeArena(aoi ArenaOfInterest) bool {
	var (
		now   time.Time
		next  time.Time
		last  time.Time
		delay time.Duration
	)

	// need to compare current time with the next time it should run
	// and update the values as required

	now = time.Now()
	delay = time.Duration{time.Minute * time.Duration(aoi.ScrapeDelayMin)}
	last = time.Unix(aoi.LastRun, 0)
	next = last.Add(delay)

	if now.After(next) {
		return true
	}
	return false
}

func within(aoi ArenaOfInterest) bool {
	var (
		current Clock
		now     time.Time
	)
	now = time.Now()
	current = Clock{}
	current.Hour, current.Minute, current.Second = now.Clock()
	afterTime

}
