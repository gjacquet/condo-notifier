package processor

import "github.com/mmcdole/gofeed"
import "github.com/gjacquet/condo-notifier/condo"

func ParseRss(url string) ([]*condo.CondoAd, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}

	ads := make([]*condo.CondoAd, len(feed.Items))
	for i, item := range feed.Items {
		ads[i] = &condo.CondoAd{
			Title: item.Title,
		}
	}

	return ads, nil
}
