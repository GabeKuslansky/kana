package anki

import (
	"bytes"
	"encoding/json"
	"errors"
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

type DeckNamesAndIdsResponse struct {
	Result map[string]int `json:"result"`
}

type CreateDeckParams struct {
	Deck string `json:"deck"`
}

type CreateDeckResponse struct {
	Id int `json:"result"`
}

type DeckStatsResponse struct {
	Result map[string]DeckStatsResult
	Error  interface{}
}

type DeckStatsResult struct {
	Deck_Id       uint   `json:"deck_id"`
	Name          string `json:"name"`
	New_Count     uint   `json:"new_count"`
	Learn_Count   uint   `json:"learn_count"`
	Review_Count  uint   `json:"review_count"`
	Total_In_Deck uint   `json:"total_in_deck"`
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

	var response DeckNamesAndIdsResponse
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}

func GetDeckStats(names []string) (map[string]DeckStatsResult, error) {
	params := map[string]interface{}{
		"decks": names,
	}

	jsonBytes, err := getJsonBytes("getDeckStats", params)
	if err != nil {
		return nil, err
	}

	respBytes, err := Request(jsonBytes)
	if err != nil {
		return nil, err
	}

	var response DeckStatsResponse
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

	var response CreateDeckResponse
	err = json.Unmarshal(resBytes, &response)
	if err != nil {
		return 0, err
	}

	return response.Id, nil
}

type AddCardResponse struct {
	Id  int
	Err string
}

func AddCard(front string, back string, deck string) (int, error) {
	params := map[string]interface{}{
		"note": map[string]interface{}{
			"deckName":  deck,
			"modelName": "Basic",
			"fields": map[string]interface{}{
				"Front": front,
				"Back":  back,
			},
		},
	}
	jsonBytes, err := getJsonBytes("addNote", params)
	if err != nil {
		return 0, err
	}

	resBytes, err := Request(jsonBytes)
	if err != nil {
		return 0, err
	}

	var response AddCardResponse
	err = json.Unmarshal(resBytes, &response)
	if err != nil {
		return 0, err
	}

	if response.Err != "" {
		return 0, errors.New(response.Err)
	}

	return response.Id, nil

}

type FindCardsResponse struct {
	Result []int `json:"result"`
}

func FindCardIds() ([]int, error) {
	params := map[string]interface{}{
		"query": "deck:current",
	}

	jsonBytes, err := getJsonBytes("findCards", params)
	if err != nil {
		return nil, err
	}

	resBytes, err := Request(jsonBytes)
	if err != nil {
		return nil, err
	}

	var response FindCardsResponse
	err = json.Unmarshal(resBytes, &response)
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}

type Card struct {
	Answer     string `json:"answer"`
	Question   string `json:"question"`
	DeckName   string `json:"deckName"`
	ModelName  string `json:"modelName"`
	FieldOrder int    `json:"fieldOrder"`
	Fields     struct {
		Front struct {
			Value string `json:"value"`
			Order int    `json:"order"`
		} `json:"Front"`
		Back struct {
			Value string `json:"value"`
			Order int    `json:"order"`
		} `json:"Back"`
	} `json:"fields"`
	CSS      string `json:"css"`
	CardID   int64  `json:"cardId"`
	Interval int    `json:"interval"`
	Note     int64  `json:"note"`
	Ord      int    `json:"ord"`
	Type     int    `json:"type"`
	Queue    int    `json:"queue"`
	Due      int    `json:"due"`
	Reps     int    `json:"reps"`
	Lapses   int    `json:"lapses"`
	Left     int    `json:"left"`
	Mod      int    `json:"mod"`
}

type CardsInfoResponse struct {
	Result []Card
}

func GetCardsInfo(cardIds []int) ([]Card, error) {
	params := map[string]interface{}{
		"cards": cardIds,
	}

	jsonBytes, err := getJsonBytes("cardsInfo", params)
	if err != nil {
		return nil, err
	}

	resBytes, err := Request(jsonBytes)
	if err != nil {
		return nil, err
	}

	var response CardsInfoResponse
	err = json.Unmarshal(resBytes, &response)
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}
