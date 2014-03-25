package feeder

import (
	"crypto/sha1"
	"fmt"
	"time"
)

type Item struct {
	// RSS and Shared fields
	Title       string
	Links       []*Link
	Description string
	Author      Author
	Categories  []*Category
	Comments    string
	Enclosures  []*Enclosure
	Guid        *string
	PubDate     string
	Source      *Source

	// Atom specific fields
	Id           string
	Generator    *Generator
	Contributors []string
	Content      *Content

	Extensions map[string]map[string][]Extension
}

func (i *Item) ParsedPubDate() (time.Time, error) {
	return parseTime(i.PubDate)
}

func (i *Item) Key() string {
	return i.hash(i.keyString())
}

func (i *Item) keyString() string {
	switch {
	case i.Guid != nil && len(*i.Guid) != 0:
		return *i.Guid
	case len(i.Id) != 0:
		return i.Id
	case len(i.Title) > 0 && len(i.PubDate) > 0:
		return i.Title + i.PubDate
	default:
		return i.Description
	}
}

func (i *Item) hash(keyString string) string {
	h := sha1.New()
	h.Write([]byte(keyString))
	return fmt.Sprintf("%x", h.Sum(nil))
}
