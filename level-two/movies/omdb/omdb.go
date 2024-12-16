package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Result struct {
	Search       []SearchResult `json:"Search"`
	TotalResults string   `json:"totalResults"`
	Response     string   `json:"Response"`
}

type SearchResult struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"ImdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

func Search(apiKey, title string) (Result, error) {
	v := url.Values{}

	v.Set("apikey", apiKey)
	v.Set("s", title)

	resp, err := http.Get("http://www.omdbapi.com/?" + v.Encode())
	if err != nil {
		return Result{}, fmt.Errorf("failed to make request to omdb: %w", err)
	}
	defer resp.Body.Close()

	var result Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Result{}, fmt.Errorf("failed to decode response from omdb: %w", err)
	}
	return result, nil
}