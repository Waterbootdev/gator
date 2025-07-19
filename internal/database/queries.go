package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Waterbootdev/gator/internal/feeds"
)

func (q *Queries) ScrapeFeeds(user User) error {

	feed, err := q.GetNextFeedToFetch(context.Background(), user.ID)
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

	rSSFeed, err := feeds.FetchFeed(context.Background(), feed.Url)

	if err != nil {
		return err
	}

	title := feed.Name + ":"

	for _, item := range rSSFeed.Channel.Item {
		if len(item.Title) == 0 {
			continue
		}
		fmt.Println(title, item.Title)
	}

	return nil
}
