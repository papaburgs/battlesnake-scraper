package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func Scrape(league string, snakeChan chan AreanSnake) {

	var httpClient = &http.Client{Timeout: 10 * time.Second}

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		var err error
		var s Snake

		s.UUID = e.Attr("id")
		s.Rank = e.ChildText(".arena-leaderboard-rank")
		sName := strings.Split(e.ChildText(".arena-leaderboard-name"), "\n")[0]
		s.Name = strings.TrimSpace(sName)
		s.Rating = e.ChildText(".arena-leaderboard-rating")
		s.League = league

		// create/update snake
		err = a.UpdateSnake(s)
		if err != nil {
			log.Fatal(err)
		}
		snakeChan <- s
		// skip table rows without a snake UUID (division rows)
		if len(s.UUID) != 0 {
			log.Printf("%+v - %+v - %+v - %+v - %+v\n", league, s.Rank, s.Name, s.Rating, s.UUID)
			recentGamesURL := fmt.Sprintf("https://play.battlesnake.com/arena/details/%s/", s.UUID)
			r, err := httpClient.Get(recentGamesURL)
			if err != nil {
				log.Fatal(err)
			}
			defer r.Body.Close()

			rgr := RecentGamesResponse{}
			err = json.NewDecoder(r.Body).Decode(&rgr)
			if err != nil {
				log.Fatal(err)
			}
			for _, recentGame := range rgr.RecentGames {
				recentGame.GameUUID = strings.Split(recentGame.GameURL, "/")[2]

				// check if recent game record already in database
				if !a.RecentExists(league, recentGame.GameUUID, s) {
					err := a.SaveRecentGame(s, recentGame)
					if err != nil {
						log.Fatal(err)
					}
				}

				// avoid web request if record already exists in database
				if !a.GameExists(league, recentGame.GameUUID) {
					time.Sleep(time.Second * 1) // be nice :)
					log.Printf("retrieving game info: %s - %s\n", league, recentGame.GameUUID)
					gameInfoURL := fmt.Sprintf("https://engine.battlesnake.com/games/%s", recentGame.GameUUID)
					r, err := httpClient.Get(gameInfoURL)
					if err != nil {
						log.Fatal(err)
					}
					defer r.Body.Close()

					gameInfo := GameInfo{}
					err = json.NewDecoder(r.Body).Decode(&gameInfo)
					if err != nil {
						log.Fatal(err)
					}

					err = a.SaveGame(league, gameInfo)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
			time.Sleep(time.Second * 1) // be nice :)
		}
	})

	queryURL := fmt.Sprintf("https://play.battlesnake.com/arena/%s/", league)
	c.Visit(queryURL)
}
