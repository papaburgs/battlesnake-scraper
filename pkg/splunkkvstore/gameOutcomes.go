package splunkkvstore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

const (
	collect_outcomes = "gameOutcome"
)

// GameOutcomeRecord is based off the 'engine' endpoint and
// holds the result for each snake (from 'LastFrame')
type GameOutcomeRecord struct {
	Key        string `json:"_key,omitempty"`
	GameID     string
	SnakeGSID  string
	Name       string
	DeathCause string
	DeathTurn  int
	DeathFrom  string
}

func (k KVStore) OutcomeExist(gameID string) bool {
	var (
		err     error
		address string
		//	fields string
		query  string
		output []GameOutcomeRecord
	)

	query = fmt.Sprintf("{\"GameID\": \"%s\"}", gameID)
	address = fmt.Sprintf("%s/%s?query=%s",
		k.base, collect_outcomes, url.QueryEscape(query))
	err = k.Get(address, &output)
	if err != nil {
		log.Println(err)
		return false
	}
	fmt.Println(len(output))
	return len(output) > 0
}

func (k KVStore) GetAllOutcomes() []GameOutcomeRecord {
	var (
		url   string
		err   error
		games []GameOutcomeRecord
	)
	url = fmt.Sprintf("%s/%s", k.base, collect_outcomes)
	err = k.Get(url, &games)
	if err != nil {
		log.Println(err)
	}
	return games
}

func (k KVStore) AddOutcomeGame(r GameOutcomeRecord) error {
	var (
		url  string
		data []byte
		err  error
	)
	url = fmt.Sprintf("%s/%s", k.base, collect_outcomes)
	data, err = json.Marshal(r)
	err = k.Post(url, data)
	return err
}
