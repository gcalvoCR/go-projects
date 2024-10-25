package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Quote struct {
	Quote  string
	Author string
}

const (
	domain  = "books.toscrape.com"
	baseurl = "https://books.toscrape.com"
)

func ProcessLink(s string) string {
	if strings.Contains(s, "https://") {
		return s
	}
	return fmt.Sprintf("%s/catalogue/%s", baseurl, s)
}

func RunScraper() {
	urls := []string{fmt.Sprintf("%s/catalogue/%s", baseurl, "page-1.html")}

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Status Code", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println(err, err.Error())
	})
	/*

		c.OnHTML(".quote", func(h *colly.HTMLElement) {
			div := h.DOM
			quote := div.Find(".text").Text()
			author := div.Find(".author").Text()
			quotes = append(quotes, Quote{quote, author})

		})

			c.OnHTML(".text", func(h *colly.HTMLElement) {
				fmt.Println("The text is", h.Text, "\n")
			})
	*/

	c.OnHTML(".next a", func(h *colly.HTMLElement) {
		fmt.Println("found")
		link := h.Attr("href")
		fmt.Println("href", link)
		url := ProcessLink(link)
		urls = append(urls, url)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Page scraped")
	})

	for len(urls) > 0 {
		c.Visit(urls[0])
		fmt.Println("url visited", urls[0])
		urls = urls[1:]
	}

}
