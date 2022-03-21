package splunkkvstore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

// RecentRecord is saved in splunk's mongodb
// it is in relation to the current snake
type RecentRecord struct {
	Key         string `json:"_key,omitempty"`
	GameID      string
	SnakeSAIID  string
	Turns       int
	ScoreChange string
	PointChange string
	TierChange  string
	Result      string
	Arena       string
}

const (
	collect_recentGames = "recentGames"
)

func (k KVStore) RecentExist(gameID string) bool {
	var (
		err     error
		address string
		//	fields string
		query  string
		output []RecentRecord
	)

	// fields = "game_id,result"
	query = fmt.Sprintf("{\"GameID\": \"%s\"}", gameID)
	address = fmt.Sprintf("%s/%s?query=%s",
		k.base, collect_recentGames, url.QueryEscape(query))
	err = k.Get(address, &output)
	if err != nil {
		log.Println(err)
		return false
	}
	return len(output) > 0
}

func (k KVStore) GetAllRecentGames() []RecentRecord {
	var (
		url   string
		err   error
		games []RecentRecord
	)
	url = fmt.Sprintf("%s/%s", k.base, collect_recentGames)
	err = k.Get(url, &games)
	if err != nil {
		log.Println(err)
	}
	return games
}

func (k KVStore) AddRecentGame(r RecentRecord) error {
	var (
		url  string
		data []byte
		err  error
	)
	url = fmt.Sprintf("%s/%s", k.base, collect_recentGames)
	data, err = json.Marshal(r)
	err = k.Post(url, data)
	return err
}
