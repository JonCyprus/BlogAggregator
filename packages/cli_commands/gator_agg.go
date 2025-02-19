package cli_commands

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/JonCyprus/BlogAggregator/internal/database"
	"github.com/JonCyprus/BlogAggregator/packages/feed"
	"github.com/lib/pq"
	"time"
)

func handleAgg(state *State, cmd Command) error {
	// Checking for correct args
	args := cmd.GetArgs()
	if len(args) != 1 {
		fmt.Println("Usage: agg <time_between_reqs>")
		return ErrInvalidArgument
	}
	// Get the time duration
	timeBtwnReq, err := time.ParseDuration(args[0])
	if err != nil {
		fmt.Println("Invalid time duration")
		return ErrInvalidArgument
	}
	fmt.Printf("Collecting feeds every %v\n", timeBtwnReq)

	// Create a ticker
	ticker := time.NewTicker(timeBtwnReq)
	for ; true; <-ticker.C {
		err = scrapeFeeds(state)
		if err != nil {
			fmt.Println("Error while scraping feeds")
			return err
		}
	}
	return nil
}

func scrapeFeeds(state *State) error {
	queries := state.GetDB()
	nextFeedToGet, err := queries.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("Error while fetching next feed to get")
		return err
	}
	markedFeed, err := queries.MarkFeedFetched(context.Background(), nextFeedToGet.ID)
	if err != nil {
		fmt.Println("Error while marking feed to get")
		return err
	}
	feedData, err := feed.FetchFeed(context.Background(), markedFeed.Url)
	if err != nil {
		fmt.Println("Error while fetching feed")
		return err
	}
	for _, item := range feedData.Channel.Item {
		params, err := InitCreatePostParams(state, markedFeed.Url, item)
		if err != nil {
			return err
		}
		_, err = queries.CreatePost(context.Background(), params) // could take post for debugging
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			continue
		} else {
			return err
		}
	}
	return nil
}

func InitCreatePostParams(state *State, url string, item feed.RSSItem) (database.CreatePostParams, error) {
	queries := state.GetDB()
	feedData, err := queries.GetFeedByURL(context.Background(), url)
	if err != nil {
		fmt.Println("Error while fetching feed data")
		return database.CreatePostParams{}, err
	}

	return database.CreatePostParams{
		Url:         item.Link,
		Description: sql.NullString{String: item.Description, Valid: true},
		PublishedAt: sql.NullString{String: item.PubDate, Valid: true},
		FeedID:      feedData.ID,
	}, nil
}
