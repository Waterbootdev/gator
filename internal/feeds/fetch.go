package feeds

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "gator")
	req.Header.Add("Content-Type", "application/xml")

	client := &http.Client{
		Timeout: time.Second * 10, // Timeout each requests
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	feed := &RSSFeed{}

	err = xml.Unmarshal(body, feed)

	if err != nil {
		return nil, err
	}

	feed.unescapeTitelDescription()

	return feed, nil
}
