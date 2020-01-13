package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	LiveRetryTimeout      int       `json:"live_retry_timeout"`
	Lang                  string    `json:"preferred_language"`
	CheckUpdate           bool      `json:"check_updates"`
	SaveLogs              bool      `json:"save_logs"`
	LogLocation           string    `json:"log_location"`
	DownloadLocation      string    `json:"download_location"`
	CustomPlaybackOptions []command `json:"custom_playback_options"`
	HorizontalLayout      bool      `json:"horizontal_layout"`
	Theme                 theme     `json:"theme"`
}

type theme struct {
	BackgroundColor     string `json:"background_color"`
	BorderColor         string `json:"border_color"`
	CategoryNodeColor   string `json:"category_node_color"`
	FolderNodeColor     string `json:"folder_node_color"`
	ItemNodeColor       string `json:"item_node_color"`
	ActionNodeColor     string `json:"action_node_color"`
	LoadingColor        string `json:"loading_color"`
	LiveColor           string `json:"live_color"`
	UpdateColor         string `json:"update_color"`
	NoContentColor      string `json:"no_content_color"`
	InfoColor           string `json:"info_color"`
	ErrorColor          string `json:"error_color"`
	TerminalAccentColor string `json:"terminal_accent_color"`
	TerminalTextColor   string `json:"terminal_text_color"`
}

func loadConfig() (cfg config, err error) {
	path, err := getConfigPath()
	if err != nil {
		return
	}

	if _, err = os.Stat(path + "config.json"); os.IsNotExist(err) {
		cfg.LiveRetryTimeout = 60
		cfg.Lang = "en"
		cfg.CheckUpdate = true
		cfg.SaveLogs = true
		err = cfg.save()
		return
	}

	var data []byte
	data, err = ioutil.ReadFile(path + "config.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &cfg)
	return
}

func (cfg config) save() error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(&cfg, "", "\t")
	if err != nil {
		return fmt.Errorf("error marshaling config: %v", err)
	}

	err = ioutil.WriteFile(path+"config.json", data, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error saving config: %v", err)
	}
	return nil
}
