package itunes

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type ItunesService struct{}

func NewItunesService() *ItunesService {
	return &ItunesService{}
}

type SearchResponse struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		WrapperType            string    `json:"wrapperType"`
		ArtistID               int       `json:"artistId,omitempty"`
		CollectionID           int       `json:"collectionId,omitempty"`
		ArtistName             string    `json:"artistName"`
		CollectionName         string    `json:"collectionName,omitempty"`
		CollectionCensoredName string    `json:"collectionCensoredName,omitempty"`
		ArtistViewURL          string    `json:"artistViewUrl,omitempty"`
		CollectionViewURL      string    `json:"collectionViewUrl,omitempty"`
		ArtworkURL60           string    `json:"artworkUrl60"`
		ArtworkURL100          string    `json:"artworkUrl100"`
		CollectionPrice        float64   `json:"collectionPrice,omitempty"`
		CollectionExplicitness string    `json:"collectionExplicitness"`
		TrackCount             int       `json:"trackCount,omitempty"`
		Copyright              string    `json:"copyright,omitempty"`
		Country                string    `json:"country"`
		Currency               string    `json:"currency"`
		ReleaseDate            time.Time `json:"releaseDate"`
		PrimaryGenreName       string    `json:"primaryGenreName"`
		PreviewURL             string    `json:"previewUrl,omitempty"`
		Description            string    `json:"description,omitempty"`
		Kind                   string    `json:"kind,omitempty"`
		TrackID                int       `json:"trackId,omitempty"`
		TrackName              string    `json:"trackName,omitempty"`
		TrackCensoredName      string    `json:"trackCensoredName,omitempty"`
		TrackViewURL           string    `json:"trackViewUrl,omitempty"`
		ArtworkURL30           string    `json:"artworkUrl30,omitempty"`
		TrackPrice             float64   `json:"trackPrice,omitempty"`
		TrackExplicitness      string    `json:"trackExplicitness,omitempty"`
		DiscCount              int       `json:"discCount,omitempty"`
		DiscNumber             int       `json:"discNumber,omitempty"`
		TrackNumber            int       `json:"trackNumber,omitempty"`
		TrackTimeMillis        int       `json:"trackTimeMillis,omitempty"`
		IsStreamable           bool      `json:"isStreamable,omitempty"`
		CollectionArtistName   string    `json:"collectionArtistName,omitempty"`
		ContentAdvisoryRating  string    `json:"contentAdvisoryRating,omitempty"`
		TrackRentalPrice       float64   `json:"trackRentalPrice,omitempty"`
		CollectionHdPrice      float64   `json:"collectionHdPrice,omitempty"`
		TrackHdPrice           float64   `json:"trackHdPrice,omitempty"`
		TrackHdRentalPrice     float64   `json:"trackHdRentalPrice,omitempty"`
		ShortDescription       string    `json:"shortDescription,omitempty"`
		LongDescription        string    `json:"longDescription,omitempty"`
		FeedURL                string    `json:"feedUrl,omitempty"`
		ArtworkURL600          string    `json:"artworkUrl600,omitempty"`
		GenreIds               []string  `json:"genreIds,omitempty"`
		Genres                 []string  `json:"genres,omitempty"`
	} `json:"results"`
}

func (itunesService *ItunesService) search(term string) (SearchResponse, error) {
	client := http.Client{Timeout: 10 * time.Second}

	url := url.URL{Scheme: "https", Host: "itunes.apple.com", Path: "search"}
	query := url.Query()
	query.Set("term", term)
	query.Set("media", "music")
	url.RawQuery = query.Encode()

	res, err := client.Get(url.String())
	if err != nil {
		return SearchResponse{}, err
	}

	searchResponse := SearchResponse{}

	err = json.NewDecoder(res.Body).Decode(searchResponse)
	if err != nil {
		return SearchResponse{}, err
	}

	return searchResponse, nil
}
