package main

import (
	_ "embed"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/papaburgs/battlesnake-scraper/pkg/scraper"
	"github.com/papaburgs/battlesnake-scraper/pkg/splunkkvstore"
)

/* Configuration
config is done in a really basic way.
make a text file that has 4 lines in it like:
```
hostname
8089
really-long-token
splunkapp
```
Those will be read into 'configContents' below and then
are split out to make the splunkkvstore object
*/

//go:embed config.txt
var configContents string

//go:embed gitc.txt
var gitCommit string

func main() {
	var (
		snakes      []scraper.ArenaSnake
		kvstore     *splunkkvstore.KVStore
		err         error
		configArray []string
	)
	log.Printf("starting scraper version: %s", gitCommit)
	configArray = strings.Split(configContents, "\n")
	kvstore = splunkkvstore.NewKVStore(
		configArray[0], // host
		configArray[1], // port
		configArray[2], // token
		configArray[3], // app
		"nobody",       //user
	)
	arenas := []string{"global-wrapped"}
	for _, arena := range arenas {
		log.Print("Scraping")
		snakes = scraper.Scrape(arena)

		beNice(10000)
		for _, snake := range snakes {
			if snake.Name == "Mollywobbles" {
				beNice(10000)
				log.Print("getting recent games")
				recent := scraper.GetRecentGames(snake.SAIID)
				for _, r := range recent {
					if !kvstore.RecentExist(r.GameID) {
						err = kvstore.AddRecentGame(response2record(r, arena))
						if err == nil {
							log.Printf("Added game %s", r.GameID)
						} else {
							log.Println(err)
						}
					} else {
						log.Printf("Game %s already exists", r.GameID)
					}

					if !kvstore.OutcomeExist(r.GameID) {
						beNice(10000)
						gameInfo := scraper.GetLastFrame(r.GameID)
						snakeOutcomes := gameInfo2splunk(gameInfo)
						for _, snakeOutcome := range snakeOutcomes {
							err = kvstore.AddOutcomeGame(snakeOutcome)
							if err == nil {
								log.Printf("Added game %s", r.GameID)
							} else {
								log.Println(err)
							}
						}
					} else {
						log.Printf("Game %s already exists", r.GameID)
					}

				}
			}
		}

	}
}

func response2record(s scraper.RecentGame, arena string) splunkkvstore.RecentRecord {
	c := splunkkvstore.RecentRecord{
		GameID:      s.GameID,
		SnakeSAIID:  s.SnakeID,
		Turns:       s.Turns,
		ScoreChange: s.ScoreChange,
		PointChange: s.PointChange,
		TierChange:  s.TierChange,
		Result:      s.Result,
		Arena:       arena,
	}
	return c
}

func gameInfo2splunk(g scraper.GameInfo) []splunkkvstore.GameOutcomeRecord {
	var or []splunkkvstore.GameOutcomeRecord
	var r splunkkvstore.GameOutcomeRecord
	or = []splunkkvstore.GameOutcomeRecord{}

	for _, snake := range g.LastFrame.Snakes {
		r = splunkkvstore.GameOutcomeRecord{
			GameID:     g.Game.ID,
			SnakeGSID:  snake.ID,
			Name:       snake.Name,
			DeathCause: snake.Death.Cause,
			DeathTurn:  snake.Death.Turn,
			DeathFrom:  snake.Death.EliminatedBy,
		}
		or = append(or, r)
	}

	return or
}

func beNice(ms int) {
	// set random number of milliseconds
	rand.Seed(time.Now().Unix())
	rms := rand.Intn(ms)
	time.Sleep(time.Duration(rms) * time.Millisecond)
}
