package itunes

import (
	"encoding/xml"
	"time"
)

type SearchItem struct {
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
}

type SearchResponse struct {
	ResultCount int          `json:"resultCount"`
	Results     []SearchItem `json:"results"`
}

type ChannelItem struct {
	Text        string `xml:",chardata"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Guid        struct {
		Text        string `xml:",chardata"`
		IsPermaLink string `xml:"isPermaLink,attr"`
	} `xml:"guid"`
	Creator   string `xml:"creator"`
	PubDate   string `xml:"pubDate"`
	Enclosure struct {
		Text   string `xml:",chardata"`
		URL    string `xml:"url,attr"`
		Length string `xml:"length,attr"`
		Type   string `xml:"type,attr"`
	} `xml:"enclosure"`
	Summary  string `xml:"summary"`
	Explicit string `xml:"explicit"`
	Duration string `xml:"duration"`
	Image    struct {
		Text string `xml:",chardata"`
		Href string `xml:"href,attr"`
	} `xml:"image"`
	Season      string `xml:"season"`
	Episode     string `xml:"episode"`
	EpisodeType string `xml:"episodeType"`
}

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Dc      string   `xml:"dc,attr"`
	Content string   `xml:"content,attr"`
	Atom    string   `xml:"atom,attr"`
	Version string   `xml:"version,attr"`
	Itunes  string   `xml:"itunes,attr"`
	Anchor  string   `xml:"anchor,attr"`
	Channel struct {
		Text        string `xml:",chardata"`
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Link        struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Image struct {
			Text  string `xml:",chardata"`
			Href  string `xml:"href,attr"`
			URL   string `xml:"url"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"image"`
		Generator     string `xml:"generator"`
		LastBuildDate string `xml:"lastBuildDate"`
		Author        string `xml:"author"`
		Copyright     string `xml:"copyright"`
		Language      string `xml:"language"`
		Summary       string `xml:"summary"`
		Type          string `xml:"type"`
		Owner         struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name"`
			Email string `xml:"email"`
		} `xml:"owner"`
		Explicit string `xml:"explicit"`
		Category struct {
			Text     string `xml:",chardata"`
			AttrText string `xml:"text,attr"`
			Category struct {
				Text     string `xml:",chardata"`
				AttrText string `xml:"text,attr"`
			} `xml:"category"`
		} `xml:"category"`
		Item []ChannelItem `xml:"item"`
	} `xml:"channel"`
}
