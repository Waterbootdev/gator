package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/Waterbootdev/gator/internal/feeds"
	"github.com/google/uuid"
)

const layout string = "Mon, 02 Jan 2006 15:04:05 -0700"

func sqlTime(t string) sql.NullTime {
	parsedTime, err := time.Parse(layout, t)

	if err != nil {
		return sql.NullTime{
			Time:  parsedTime,
			Valid: false,
		}
	}

	return sql.NullTime{
		Time:  parsedTime,
		Valid: true,
	}
}

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

	currentTime := time.Now()

	for _, item := range rSSFeed.Channel.Item {

		q.CreatePost(context.Background(), CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: sqlTime(item.PubDate),
			FeedID:      feed.ID,
		})
	}

	return nil
}
