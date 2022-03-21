package splunkkvstore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

// SnakeRecord is saved in splunk's mongodb
// it is in relation to the current snake
type SnakeRecord struct {
	Key        string `json:"_key,omitempty"`
	SnakeSAIID string
	SnakeGSID  string
	SnakeName  string
	Arena      string
}

const (
	collect_snakeID = "snakeID"
)

func (k KVStore) GetSnake(s string) []SnakeRecord {
	var (
		err     error
		address string
		//	fields string
		query  string
		output []SnakeRecord
	)

	query = fmt.Sprintf("{\"SnakeSAIID\": \"%s\"}", s)
	address = fmt.Sprintf("%s/%s?query=%s",
		k.base, collect_snakeID, url.QueryEscape(query))
	err = k.Get(address, &output)
	if err != nil {
		log.Println(err)
	}
	return output
}

func (k KVStore) SAIIDExist(s string) bool {
	return len(k.GetSnake(s)) > 0
}

func (k KVStore) GetAllSnakes() []SnakeRecord {
	var (
		url   string
		err   error
		games []SnakeRecord
	)
	url = fmt.Sprintf("%s/%s", k.base, collect_snakeID)
	err = k.Get(url, &games)
	if err != nil {
		log.Println(err)
	}
	return games
}

func (k KVStore) AddSnakeRecord(r SnakeRecord) error {
	var (
		url  string
		data []byte
		err  error
	)
	url = fmt.Sprintf("%s/%s", k.base, collect_snakeID)
	data, err = json.Marshal(r)
	err = k.Post(url, data)
	return err
}
