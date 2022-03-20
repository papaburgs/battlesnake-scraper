package splunkkvstore

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
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

func NewKVStore(host, port, token, app, user string) *KVStore {
	k := KVStore{
		token: token,
		port:  port,
		host:  host,
		app:   app,
		user:  user,
	}
	k.base = fmt.Sprintf("https://%s:%s/servicesNS/nobody/%s/storage/collections/data/",
		k.host, k.port, k.app)
	return &k
}

func (k KVStore) Get(url string, output interface{}) error {
	var (
		client  *http.Client
		resp    *http.Response
		err     error
		content io.Reader
		bearer  string
	)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr, Timeout: 10 * time.Second}

	bearer = "Bearer " + k.token
	content = bytes.NewReader([]byte(`{"output_mode": "json"}`))
	req, err := http.NewRequest("GET", url, content)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		log.Println("Status: ", resp.Status)
		return errors.New(resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(output)
	if err != nil {
		return err
	}
	return nil
}

func (k KVStore) Post(url string, data []byte) error {
	var (
		client  *http.Client
		resp    *http.Response
		err     error
		content io.Reader
		bearer  string
	)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr, Timeout: 10 * time.Second}

	bearer = "Bearer " + k.token
	content = bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, content)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		log.Println("Status: ", resp.Status)
		return errors.New(resp.Status)
	}
	return nil
}
