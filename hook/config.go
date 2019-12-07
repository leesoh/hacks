package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

type config struct {
	HookURL string `json:"hook_url"`
}

func (c *config) Load() error {
	url, ok := os.LookupEnv("SLACK_HOOK")
	if ok {
		c.HookURL = url
	}
	home := os.Getenv("HOME")
	configFile := path.Join(home, ".config", "hook", ".hook.json")
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c)
	return nil
}
