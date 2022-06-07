package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"itunesJsonXmlApis/graph"
	"itunesJsonXmlApis/graph/generated"
	"log"
	"net/http"
	"os"

	"itunesJsonXmlApis/itunes"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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
