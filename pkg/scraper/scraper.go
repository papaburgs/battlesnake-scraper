package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func Scrape(league string) []ArenaSnake {
	snakes := []ArenaSnake{}

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		// var err error
		var s ArenaSnake

		s.SAIID = e.Attr("id")
		s.Rank = e.ChildText(".arena-leaderboard-rank")
		sName := strings.Split(e.ChildText(".arena-leaderboard-name"), "\n")[0]
		s.Name = strings.TrimSpace(sName)
		s.Rating = e.ChildText(".arena-leaderboard-rating")
		s.League = league

		if len(s.SAIID) != 0 {
			snakes = append(snakes, s)
		}
	})

	queryURL := fmt.Sprintf("https://play.battlesnake.com/arena/%s/", league)
	c.Visit(queryURL)
	return snakes
}

// GetRecentGames returns a list of 20 recent games based on a instance ID
func GetRecentGames(saiid string) []RecentGame {
	var (
		recentGamesURL string
		g              RecentGame
		err            error
		recentGames    []RecentGame
		httpClient     *http.Client
		resp           *http.Response
		rgr            RecentGamesResponse
	)
	recentGames = []RecentGame{}
	httpClient = &http.Client{Timeout: 10 * time.Second}
	recentGamesURL = fmt.Sprintf("https://play.battlesnake.com/arena/details/%s/", saiid)
	resp, err = httpClient.Get(recentGamesURL)
	if err != nil {
		log.Printf("%s", err)
		return recentGames
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&rgr)
	if err != nil {
		log.Printf("%s", err)
	}
	for _, recentGame := range rgr.RecentGames {
		g = recentGame
		g.GameID = strings.Split(recentGame.GameURL, "/")[2]
		g.SnakeID = saiid
		recentGames = append(recentGames, g)
	}
	return recentGames
}

func GetLastFrame(gid string) GameInfo {
	var (
		err        error
		httpClient *http.Client
		resp       *http.Response
		gameInfo   GameInfo
		payload    []byte
	)
	if len(gid) == 0 {
		log.Print("No GID provided")
		return GameInfo{}
	}
	log.Printf("retrieving game info: %s\n", gid)
	gameInfoURL := fmt.Sprintf("https://engine.battlesnake.com/games/%s", gid)
	httpClient = &http.Client{Timeout: 10 * time.Second}
	resp, err = httpClient.Get(gameInfoURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if payload, err = io.ReadAll(resp.Body); err != nil {
		log.Printf("ERROR: Failed to read body, %s", err)
		return gameInfo
	}

	if err = json.Unmarshal(payload, &gameInfo); err != nil {
		log.Printf("ERROR: Failed to decode start json, %s", err)
		return gameInfo
	}

	return gameInfo

}
