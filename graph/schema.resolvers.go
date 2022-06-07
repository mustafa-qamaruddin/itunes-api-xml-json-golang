package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"itunesJsonXmlApis/graph/generated"
	"itunesJsonXmlApis/graph/model"
	"itunesJsonXmlApis/itunes"
	"log"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Search(ctx context.Context, term string, media *string, entity *string, attribute *string) ([]*model.SearchItem, error) {
	itunesService := itunes.NewItunesService()

	res, err := itunesService.Search(term, media, entity, attribute, 200)
	if err != nil {
		log.Fatal(err.Error())
	}

	var ret []*model.SearchItem

	for _, item := range res.Results {

		var feeds []*model.ChannelItem
		if item.FeedURL != "" {
			feed, err := itunesService.FetchFeed(item.FeedURL)
			if err != nil {
				log.Fatal(err.Error())
			}
			for _, feedItem := range feed.Channel.Item {
				feeds = append(feeds, &model.ChannelItem{
					Title:       feedItem.Title,
					Duration:    &feedItem.Duration,
					Description: &feedItem.Description,
					URL:         feedItem.Link,
				})
			}
		}

		searchItem := &model.SearchItem{
			ID:         item.CollectionID,
			Artist:     item.ArtistName,
			Track:      item.TrackName,
			URL:        item.TrackViewURL,
			Country:    &item.Country,
			Collection: &item.CollectionName,
			FeedURL:    &item.FeedURL,
			Length:     &item.TrackTimeMillis,
			Feeds:      feeds,
		}

		ret = append(ret, searchItem)
	}
	return ret, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
