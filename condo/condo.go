package condo

import "net/url"
import "time"
import "golang.org/x/text/currency"
import "github.com/golang/geo/s2"

type CondoAd struct {
	Title         string          `json:"title,omitempty"`
	Link          url.URL         `json:"link,omitempty"`
	Description   string          `json:"description,omitempty"`
	PictureLink   url.URL         `json:"pictureLink,omitempty"`
	PublishedDate time.Time       `json:"publishedDate,omitempty"`
	Price         currency.Amount `json:"price,omitempty"`
	Location      s2.LatLng       `json:"location,omitempty"`
}
