package kvstore

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type KVStore struct {
	token string
	host  string
	port  string
	base  string
	user  string
	app   string
}

const (
	collect_recentGames = "recentGames"
	collect_outcomes    = "gameOutcomes"
)

func NewKVStore(host, port, token, app, user string) *KVStore {
	k := KVStore{
		token: token,
		port:  port,
		host:  host,
		app:   app,
		user:  user,
		base:  "/servicesNS",
	}
	return &k
}

func (k *KVStore) UpdateApp(app string) {
	k.app = app
}

// ListCollections may not be used
func (k KVStore) ListCollections() {
	var (
		url         string
		collections string
	)
	url = fmt.Sprintf("https://%s:%s/%s/%s/storage/collections/config",
		k.host, k.port, k.base, k.user, k.app)
}

func (k KVStore) RecentExist(gameID string) bool {
	var (
		url    string
		fields string
		query  string
		output RecentRecord
	)

	url = fmt.Sprintf("https://%s:%s/%s/%s/storage/collections/data/%s",
		k.host, k.port, k.base, k.user, k.app, collect_recentGames)
	fields = "fields=game_id,result"
	query = fmt.Sprintf("query={\"game_id\": \"%s\"}", gameID)
	k.Get(url, &output)
	return len(output) > 0
}

func (k KVStore) AddRecentGame(r RecentRecord) error {
	var (
		url  string
		data string
	)
	url = fmt.Sprintf("https://%s:%s/%s/%s/storage/collections/data/%s",
		k.host, k.port, k.base, k.user, k.app, collect_recentGames)

	k.Post(url, data)
}

func (k KVStore) Get(url string, output interface{}) error {
	var (
		httpClient *http.Client
		resp       *http.Response
		err        error
	)
	httpClient = &http.Client{Timeout: 10 * time.Second}
	resp, err = httpClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(output)
	if err != nil {
		return err
	}
	return nil
}

func (k KVStore) Post(url string, data string) error {
	var (
		httpClient *http.Client
		resp       *http.Response
		err        error
	)
	httpClient = &http.Client{Timeout: 10 * time.Second}
	resp, err = httpClient.Post(url, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(output)
	if err != nil {
		return err
	}
	return nil
}

// endpoint storage/collections/config
// basic call

/*
curl -H "Authorization: Bearer eyJraWQiOiJzcGx1bmsuc2VjcmV0IiwiYWxnIjoiSFM1MTIiLCJ2ZXIiOiJ2MiIsInR0eXAiOiJzdGF0aWMifQ.eyJpc3MiOiJkYXZpZCBmcm9tIGFyY2hlbWVkZXMubG9jYWwiLCJzdWIiOiJkYXZpZCIsImF1ZCI6Imxvb2t1cCBtYW5pcHVsYXRpb24iLCJpZHAiOiJTcGx1bmsiLCJqdGkiOiJkYmY1NTU2MDhiNWVkZDVjOWQzNmE1NTBjYWEwM2EwNzU2YmRmMDdmYjBkNDkyMmViMGRmODM0NDgwMDIzODk2IiwiaWF0IjoxNjQ3MTM2MzYwLCJleHAiOjE2NDk3MjgzNjAsIm5iciI6MTY0NzEzNjM2MH0.f1lQN0z5tAkx0H2nS_ITiUuCeGXIR8UyZoEFN_7zLg6JLTy4p25OcekExlkN6jDso7zbQl1Yi5Od9PaXWVC9LA" -k --get -d output_mode=json https://archemedes.home.32ohsix.ca:8089/servicesNS/nobody/battlesnake/storage/collections/data/recentGames
*/
