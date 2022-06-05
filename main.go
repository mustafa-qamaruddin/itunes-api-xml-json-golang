package itunes_api_xml_json_golang

import (
	"log"

	"itunesClients/itunes"
)

func main() {
	itunesService := itunes.NewItunesService()

	res, err := itunesService.Search("mercedes sosa")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range res.Results {
		log.Println("---")
		log.Println("Artist: ", item.ArtistName)
		log.Println("Country: ", item.Country)
		log.Println("ID: ", item.CollectionID)
		log.Println("Collection: ", item.CollectionName)
		log.Println("Track: ", item.TrackName)
		log.Println("URL: ", item.TrackViewURL)
		log.Println("---")
	}
}
