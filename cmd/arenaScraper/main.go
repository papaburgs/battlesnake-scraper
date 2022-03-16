package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/papaburgs/battlesnake-scraper/pkg/scraper"
)

func main() {
	var (
		snakes []scraper.ArenaSnake
	)

	log.Print("Scraping")
	snakes = scraper.Scrape("spring-league-2022")

	beNice(10000)
	for _, snake := range snakes {
		if snake.Name == "Mollywobbles" {
			beNice(10000)
			log.Print("getting recent games")
			recent := scraper.GetRecentGames(snake.SAIID)
			fmt.Println(recent)
			beNice(10000)
			fmt.Println(scraper.GetLastFrame(recent[1].GameID))
		}
	}
}

func beNice(ms int) {
	// set random number of milliseconds
	rand.Seed(time.Now().Unix())
	rms := rand.Intn(ms)
	time.Sleep(time.Duration(rms) * time.Millisecond)
}
