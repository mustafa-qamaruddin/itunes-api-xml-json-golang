package main

import (
	"log"

	"itunesJsonXmlApis/itunes"
)

func main() {
	itunesService := itunes.NewItunesService()

	res, err := itunesService.Search("mercedes sosa", "music", "song", "artistTerm", 200)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, item := range res.Results {
		displayItem(item)
	}

	res, err = itunesService.Search("mercedes sosa", "podcast", "podcast", "keywordsTerm", 200)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, item := range res.Results {
		displayItem(item)
		if item.FeedURL != "" {
			feed, err := itunesService.FetchFeed(item.FeedURL)
			if err != nil {
				log.Fatal(err.Error())
			}
			for _, feedItem := range feed.Channel.Item {
				displayChannelItem(feedItem)
			}
		}
	}
}

func displayChannelItem(feedItem itunes.ChannelItem) {
	log.Println("* * * * *")
	log.Println("Title: ", feedItem.Title)
	log.Println("Duration: ", feedItem.Duration)
	log.Println("Description: ", feedItem.Description)
	log.Println("URL: ", feedItem.Link)
	log.Println("* * * * *")
}

func displayItem(item itunes.SearchItem) {
	log.Println("---")
	log.Println("Artist: ", item.ArtistName)
	log.Println("Country: ", item.Country)
	log.Println("ID: ", item.CollectionID)
	log.Println("Collection: ", item.CollectionName)
	log.Println("Track: ", item.TrackName)
	log.Println("URL: ", item.TrackViewURL)
	log.Println("XML Feed: ", item.FeedURL)
	log.Println("Length: ", item.TrackTimeMillis*1000)
	log.Println("---")
}
