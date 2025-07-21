package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Set("User-Agent", "gator")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, nil
	}

	rssRes := RSSFeed{}
	err = xml.Unmarshal(data, &rssRes)
	if err != nil {
		return &RSSFeed{}, err
	}

	rssRes.Channel.Title = html.UnescapeString(rssRes.Channel.Title)
	rssRes.Channel.Description = html.UnescapeString(rssRes.Channel.Description)

	for i := range rssRes.Channel.Item {
		rssRes.Channel.Item[i].Title = html.UnescapeString(rssRes.Channel.Item[i].Title)
		rssRes.Channel.Item[i].Description = html.UnescapeString(rssRes.Channel.Item[i].Description)
	}

	return &rssRes, nil
}
