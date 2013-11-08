package rssfeed

import (
	"log"
	"testing"
)

var rssList = []string{
	//"http://www.36kr.com/feed",
	"http://the5fire.com/feed",
	"http://www.ifanr.com/feed",
}

func TestFeed(t *testing.T) {
	f := Feed{}
	for _, url := range rssList {
		f.Fetch(url)
		log.Println("Channel.Title:", f.Channel.Title)
		log.Println("Channel.Description:", f.Channel.Description)
		log.Println("Channel.Link:", f.Channel.Link)
		log.Println("Channel.PubDate:", f.Channel.PubDate)

		for i, item := range f.Channel.Items {
			log.Println("item.", i)
			log.Println("------>.Item.Title:", item.Title)
			log.Println("------>.Item.Author:", item.Author)
			log.Println("------>.Item.Description:", item.Description)
			//log.Println("------>.Item.Content:", item.Content)
			log.Println("------>.Item.Guid:", item.Guid)
			log.Println("------>.Item.Link:", item.Link)
			log.Println("------>.Item.PubDate:", item.PubDate)
		}
	}
}
