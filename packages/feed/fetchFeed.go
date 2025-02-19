package feed

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Create http request with context
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		fmt.Println("Error generating request w/ context")
		return nil, fmt.Errorf("error generating request w/ context: %w", err)
	}
	// User a header to identify program to server
	req.Header.Set("User-Agent", "gator")

	// Create client and carry out request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching feed: %w", err)
	}
	defer resp.Body.Close()

	// Read the XML response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}

	// Unmarshal the response
	var feed RSSFeed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling body: %w", err)
	}

	// Unescape (decode) HTML entities in title and description
	// In the channel
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	// In every item
	for i, _ := range feed.Channel.Item {
		newTitle := html.UnescapeString(feed.Channel.Item[i].Title)
		newDescription := html.UnescapeString(feed.Channel.Item[i].Description)
		feed.Channel.Item[i].Title = newTitle
		feed.Channel.Item[i].Description = newDescription
	}

	// Return the result
	return &feed, nil
}
