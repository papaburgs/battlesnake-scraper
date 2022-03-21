package mongocfg

type MongoConfig struct {
	SnakeNameRegexes []string
	Arenas           []ArenaOfInterest
}

type ArenaOfInterest struct {
	Name           string
	SleepMin int
	StartTime Clock
	ScrapeFor int
}

type Clock struct {
	Hour   int
	Minute int
	Second int
}


/* 
say before is 7,0 and after is 23,30

it is 19,23
19 < 23 so it is not after
19 > 7 so it not before

23,20
23 >= 23, 

if a.Hour > c.Hour then after
if a.Hour = c.Hour && a.Minute < c.Minute then after
if a.Hour = c.Hour && a.Minute = c.Minute && a.Second < c.Secondthen after

if b.Hour < c.Hour then before
if b.Hour = c.Hour && b.Minute > c.Minute then before
if b.Hour = c.Hour && b.Minute = c.Minute && b.Second > c.Second then before
