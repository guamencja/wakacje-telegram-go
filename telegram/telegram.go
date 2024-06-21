package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Bot struct {
	token      string
	httpClient *http.Client
}

func New(token string) Bot {
	client := &http.Client{}

	return Bot{
		token:      token,
		httpClient: client,
	}
}

const BASE_URL = "https://api.telegram.org/bot%s/%s"

func (b *Bot) request(method string, data map[string]interface{}, result interface{}) error {
	url := fmt.Sprintf(BASE_URL, b.token, method)

	jsonStr, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := b.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var jsonStruct struct {
		Ok          bool            `json:"ok"`
		Result      json.RawMessage `json:"result"`
		Description string          `json:"description"`
	}
	json.Unmarshal(body, &jsonStruct)

	if !jsonStruct.Ok {
		return errors.New(jsonStruct.Description)
	}

	if result == nil {
		return nil
	}

	if err := json.Unmarshal(jsonStruct.Result, result); err != nil {
		return err
	}

	return nil
}
