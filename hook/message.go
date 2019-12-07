package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type Message struct {
	Text string
}

var client = &http.Client{
	Timeout: time.Second * 10,
}

func (m Message) Send(url string) error {
	b, err := json.Marshal(map[string]string{
		"text": m.Text,
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
