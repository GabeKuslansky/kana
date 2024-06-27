package anki

import (
	"bytes"
	"errors"
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

var client = &http.Client{}

type DeckNamesResponse struct {
	Result []string `json:"result"`
}

func Request(action string) ([]byte, error) {
	body := bytes.NewBuffer([]byte(DECK_NAMES_REQUEST))
	resp, err := http.Post(URL, "application/json", body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("[%s] Error querying Anki %v", action, err))
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.Bytes(), nil

}

func GetDeckNames() ([]string, error) {
	bytes, err := Request("deckNames")
	if err != nil {
		return nil, err
	}

	var response DeckNamesResponse
	json.Unmarshal(bytes, &response)

	return response.Result, nil
}
