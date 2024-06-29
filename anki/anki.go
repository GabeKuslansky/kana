package anki

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	URL     = "http://localhost:8765"
	VERSION = 6
)

type AnkiConnectRequest struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
}

type AnkiConnectRequestWithParams struct {
	Action  string                 `json:"action"`
	Version int                    `json:"version"`
	Params  map[string]interface{} `json:"params"`
}

type DeckNamesResponse struct {
	Result map[string]int `json:"result"`
	Error  interface{}    `json:"error"`
}

type CreateDeckParams struct {
	Deck string `json:"deck"`
}

type CreateDeckResult struct {
	Id int `json:"result"`
}

func getJsonBytes(action string, params interface{}) ([]byte, error) {
	var req interface{}

	if params != nil {
		req = AnkiConnectRequestWithParams{
			Action:  action,
			Version: VERSION,
			Params:  params.(map[string]interface{}),
		}
	} else {
		req = AnkiConnectRequest{
			Action:  action,
			Version: VERSION,
		}
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %v", err)
	}

	return jsonData, nil
}

func Request(jsonBytes []byte) ([]byte, error) {
	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonBytes)) // Every request to Anki-Connect is a POST
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBuf := new(bytes.Buffer)
	respBuf.ReadFrom(resp.Body)

	return respBuf.Bytes(), nil
}

func GetDeckNamesAndIds() (map[string]int, error) {
	jsonBytes, err := getJsonBytes("deckNamesAndIds", nil)
	if err != nil {
		return nil, err
	}
	respBytes, err := Request(jsonBytes)
	if err != nil {
		return nil, err
	}

	var response DeckNamesResponse
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}

func CreateDeck(name string) (int, error) {
	params := map[string]interface{}{
		"deck": name,
	}
	jsonBytes, err := getJsonBytes("createDeck", params)
	if err != nil {
		return 0, err
	}

	resBytes, err := Request(jsonBytes)
	if err != nil {
		return 0, err
	}

	var response CreateDeckResult
	err = json.Unmarshal(resBytes, &response)
	if err != nil {
		return 0, err
	}

	return response.Id, nil
}
