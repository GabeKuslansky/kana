package anki

import (
	"bytes"
	"net/http"

	// "encoding/json"
	"encoding/json"
	"fmt"
)

const (
	URL                = "http://localhost:8765"
	VERSION            = 6
	DECK_NAMES_REQUEST = `{
    "action": "deckNames",
    "version": 6
    }`
)

func formatRequest(action string) string {
	return fmt.Sprintf(`{
	"action": "%s",
	"version": %d
	}`, action, VERSION)
}

type DeckNamesResponse struct {
	Result map[string]int `json:"result"`
	Error  interface{}       `json:"error"`
}

func Request(action string) ([]byte, error) {
	body := bytes.NewBuffer([]byte(formatRequest(action)))
	resp, err := http.Post(URL, "application/json", body)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.Bytes(), nil

}

func GetDeckNamesAndIds() (map[string]int, error) {
	bytes, err := Request("deckNamesAndIds")
	if err != nil {
		return nil, err
	}

	var response DeckNamesResponse
	json.Unmarshal(bytes, &response)

	return response.Result, nil
}
