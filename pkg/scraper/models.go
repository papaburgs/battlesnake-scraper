package scraper

// ArenaSnake
type ArenaSnake struct {
	AreanaID    string        `json:"id"`
	GSID        string        `json:"gsid"`
	Rank        string        `json:"rank"`
	Name        string        `json:"name"`
	Author      string        `json:"author"`
	Rating      string        `json:"rating"`
	League      string        `json:"league"`
	RecentGames []RecentGames `json:"recent_games"`
}
