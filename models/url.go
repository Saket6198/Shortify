package models

import (
	"time"
)
type Url struct {
	Id string `json:"id`
	OriginalUrl string `json:"original_url"`
	ShortenedUrl string `json:"shortened_url"`
	CreationDate time.Time `json:"creation_date"`

}

var UrlDb = make(map[string]Url)	// small case rhega toh private hota hai identifier vrna public hoga if uppercase.........REALLY GOO!!!!!