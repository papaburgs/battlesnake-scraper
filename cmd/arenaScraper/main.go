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
		//snakes  []scraper.ArenaSnake
		kvstore *splunkkvstore.KVStore
	)
	configArray := strings.Split(configContents, "\n")
	log.Printf("starting scraper version: %s", gitCommit)
	kvstore = splunkkvstore.NewKVStore(
		configArray[0], // host
		configArray[1], // port
		configArray[2], // token
		configArray[3], // app
		"nobody",       //user
	)

	exists := kvstore.RecentExist("f3f3de63-bf57-4093-91b5-fcf84dd3f329")
	if exists {
		log.Print("exists")
	} else {
		log.Print("no exist")
	}
	//	log.Print("Scraping")
	//snakes = scraper.Scrape("spring-league-2022")

	//beNice(10000)
	//	for _, snake := range snakes {
	//		if snake.Name == "Mollywobbles" {
	//		beNice(10000)
	//			log.Print("getting recent games")
	//		recent := scraper.GetRecentGames(snake.SAIID)
	//		fmt.Println(recent[1])
	//			log.Println("trying to add game")
	/*
		err = kvstore.AddRecentGame(response2record(recent[1]))
		if err != nil {
			log.Println(err)
		}
	*/
	//		games := kvstore.GetAllRecentGames()
	//		fmt.Println(games)

	// beNice(10000)
	// fmt.Println("getting last frame")
	// fmt.Println(scraper.GetLastFrame(recent[1].GameID))
	//		}
	//	}
}

func beNice(ms int) {
	// set random number of milliseconds
	rand.Seed(time.Now().Unix())
	rms := rand.Intn(ms)
	time.Sleep(time.Duration(rms) * time.Millisecond)
}

func response2record(s scraper.RecentGame) splunkkvstore.RecentRecord {
	c := splunkkvstore.RecentRecord{
		GameID:      s.GameID,
		SnakeSAIID:  s.SnakeID,
		Turns:       s.Turns,
		ScoreChange: s.ScoreChange,
		PointChange: s.PointChange,
		TierChange:  s.TierChange,
		Result:      s.Result,
	}
	return c
}
