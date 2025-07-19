package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Waterbootdev/gator/internal/feeds"
)

func (q *Queries) scrapeFeeds() error {

	feed, err := q.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = q.MarkFeedFetched(context.Background(), MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})

	if err != nil {
		return err
	}

	RSSFeed, err := feeds.FetchFeed(context.Background(), feed.Url)

	if err != nil {
		return err
	}

	title := RSSFeed.Channel.Title + ":"

	for _, item := range RSSFeed.Channel.Item {
		fmt.Println(title, item.Title)
	}

	return nil
}
