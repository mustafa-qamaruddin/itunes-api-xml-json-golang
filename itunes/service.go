package itunes

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/url"
	"time"
)

type Service struct{}

func NewItunesService() *Service {
	return &Service{}
}

func (itunesService *Service) Search(term string, media *string, entity *string, attribute *string, limit int) (SearchResponse, error) {
	client := http.Client{Timeout: 10 * time.Second}

	url := url.URL{Scheme: "https", Host: "itunes.apple.com", Path: "search"}
	query := url.Query()
	query.Set("term", term)
	query.Set("media", *media)
	query.Set("entity", *entity)
	query.Set("attribute", *attribute)
	query.Set("limit", string(limit))
	url.RawQuery = query.Encode()

	res, err := client.Get(url.String())
	if err != nil {
		return SearchResponse{}, err
	}

	searchResponse := SearchResponse{}

	err = json.NewDecoder(res.Body).Decode(&searchResponse)
	if err != nil {
		return SearchResponse{}, err
	}

	return searchResponse, nil
}

func (itunesService *Service) FetchFeed(feedUrl string) (Rss, error) {
	res, err := http.Get(feedUrl)
	if err != nil {
		return Rss{}, err
	}
	rss := Rss{}
	err = xml.NewDecoder(res.Body).Decode(&rss)
	if err != nil {
		return Rss{}, err
	}
	return rss, nil
}
