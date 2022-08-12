package api

import (
	"encoding/json"
	"github.com/flexphere/discord-gorilla-cli/util"

	"github.com/charmbracelet/bubbles/list"
)

type Item struct {
	title, desc, url string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.desc }
func (i Item) URL() string         { return i.url }
func (i Item) FilterValue() string { return i.title }

type SearchResultLink struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type SearchResultMeeting struct {
	MeetingID int64              `json:"meetingId"`
	Elapsed   int                `json:"elapsed"`
	Timestamp int64              `json:"timestamp"`
	User      string             `json:"user"`
	Links     []SearchResultLink `json:"links"`
}

type SearchResult map[string][]SearchResultMeeting

func Search(keyword string) ([]list.Item, error) {
	url := "http://gorilla.bar38.org/search/" + keyword

	body, err := util.Fetch(url)
	if err != nil {
		panic(err)
	}

	var result SearchResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}

	var resultList []list.Item
	for _, v := range result {
		if len(v) <= 0 {
			continue
		}

		for _, mtg := range v {
			if len(mtg.Links) <= 0 {
				continue
			}

			for _, link := range mtg.Links {
				resultList = append(resultList, Item{
					title: link.Title,
					desc:  link.Description,
					url:   link.URL,
				})
			}
		}
	}

	return resultList, nil
}
