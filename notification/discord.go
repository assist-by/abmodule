package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	lib "github.com/assist-by/libStruct"
)

var (
	httpClient = &http.Client{Timeout: 10 * time.Second}
)

type DiscordMessage struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds"`
}

type Embed struct {
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	Color       int          `json:"color,omitempty"`
	Fields      []EmbedField `json:"fields,omitempty"`
	Footer      *EmbedFooter `json:"footer,omitempty"`
	Timestamp   string       `json:"timestamp,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type EmbedFooter struct {
	Text string `json:"text"`
}

func SendDiscordAlert(embed Embed, discordWebhookURL string) error {
	message := DiscordMessage{
		Embeds: []Embed{embed},
	}
	jsonPayload, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", discordWebhookURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func GetColorForDiscord(signal string) int {
	switch signal {
	case lib.SIGNAL_LONG.String():
		return 0x00FF00 // Green
	case lib.SIGNAL_SHORT.String():
		return 0xFF0000 // Red
	default:
		return 0x0000FF // Blue
	}
}
