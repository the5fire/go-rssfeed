package rssfeed

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Content     string `xml:"encoded"` // Go1.1.2版本不支持namespace
	Link        string `xml:"link"`
	Author      string `xml:"author"`
	Guid        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
}

type Channel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"linke"`
	PubDate     time.Time `xml: pubDate`
	Description string    `xml:"description"`
	Items       []*Item   `xml:"item"`
}

type Feed struct {
	Channel *Channel `xml:"channel"`
}

func (f *Feed) Fetch(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = f.Parse(body)
	return err
}

func (f *Feed) Parse(body []byte) error {
	err := xml.Unmarshal(body, &f)
	return err
}
