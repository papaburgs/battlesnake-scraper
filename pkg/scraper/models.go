package scraper

// ArenaSnake
type ArenaSnake struct {
	// Snake Arena Instance ID - start with slb_
	SAIID string `json:"id"`
	// Global Snake ID (starts with gs_
	GSID   string `json:"gsid"`
	Rank   string `json:"rank"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Rating string `json:"rating"`
	League string `json:"league"`
	//	RecentGames []RecentGames `json:"recent_games"`
}

type RecentGame struct {
	GameID      string `json:"game_id"`
	SnakeID     string `json:"snake_id"`
	GameURL     string `json:"game_url"`
	Turns       int    `json:"turns"`
	ScoreChange string `json:"score_change"`
	PointChange string `json:"point_change"`
	TierChange  string `json:"tier_change"`
	Result      string `json:"result"`
}

type RecentGamesResponse struct {
	RecentGames []RecentGame `json:"recent_games"`
}

type GameInfo struct {
	Game struct {
		ID      string `json:"ID"`
		Status  string `json:"Status"`
		Width   int    `json:"Width"`
		Height  int    `json:"Height"`
		Ruleset struct {
			DamagePerTurn     string `json:"damagePerTurn"`
			FoodSpawnChance   string `json:"foodSpawnChance"`
			MinimumFood       string `json:"minimumFood"`
			Name              string `json:"name"`
			ShrinkEveryNTurns string `json:"shrinkEveryNTurns"`
		} `json:"Ruleset"`
		SnakeTimeout int           `json:"SnakeTimeout"`
		MaxTurns     int           `json:"MaxTurns"`
		FoodSpawns   []interface{} `json:"FoodSpawns"`
		HazardSpawns []interface{} `json:"HazardSpawns"`
		Source       string        `json:"Source"`
	} `json:"Game"`
	LastFrame struct {
		Turn   int `json:"Turn"`
		Snakes []struct {
			ID   string `json:"ID"`
			Name string `json:"Name"`
			URL  string `json:"URL"`
			Body []struct {
				X int `json:"X"`
				Y int `json:"Y"`
			} `json:"Body"`
			Health       int    `json:"Health"`
			Death        Death  `json:"Death"`
			Color        string `json:"Color"`
			HeadType     string `json:"HeadType"`
			TailType     string `json:"TailType"`
			Latency      string `json:"Latency"`
			Shout        string `json:"Shout"`
			Squad        string `json:"Squad"`
			APIVersion   string `json:"APIVersion"`
			Author       string `json:"Author"`
			StatusCode   int    `json:"StatusCode"`
			Error        string `json:"Error"`
			TimingMicros struct {
				Connect   int `json:"Connect"`
				DNS       int `json:"DNS"`
				FirstByte int `json:"FirstByte"`
				Latency   int `json:"Latency"`
				TLS       int `json:"TLS"`
			} `json:"TimingMicros"`
			IsBot         bool `json:"IsBot"`
			IsEnvironment bool `json:"IsEnvironment"`
		} `json:"Snakes"`
		Food []struct {
			X int `json:"X"`
			Y int `json:"Y"`
		} `json:"Food"`
		Hazards []struct {
			X int `json:"X"`
			Y int `json:"Y"`
		} `json:"Hazards"`
	} `json:"LastFrame"`
}

type Death struct {
	Cause        string `json:"Cause"`
	Turn         int    `json:"Turn"`
	EliminatedBy string `json:"EliminatedBy"`
}
