package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gocolly/colly"
)

var url = flag.String("url", "https://telegram.org", "URL to crawl")

func main() {
	flag.Parse()
	log.Println("Initiate crawler")

	if len(*url) == 0 {
		log.Panicln(errors.New("crawler: empty url"))
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}

	webpath := path.Join(wd, "result", *url)
	file, err := os.Create(webpath)
	if err != nil {
		log.Panicln(err)
	}

	c := colly.NewCollector(colly.AllowedDomains(getDomain(*url)))
	crawlWebsite(c, *url, file)
}

func crawlWebsite(c *colly.Collector, url string, file *os.File) {
	log.Println("Start crawling", url)
	defer log.Println("Result saved at", file.Name())

	c.OnRequest(func(r *colly.Request) {
	})

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		h.Request.Visit(link)
	})

	c.OnHTML("html", func(h *colly.HTMLElement) {
		fmt.Println(h.Request.URL, url)
	})

	c.OnError(func(rs *colly.Response, err error) {
		log.Println("An error has occured, aborting")
		log.Println("Result saved at", file.Name())
		log.Panicln(err)
	})

	c.Visit(url)

}

func getDomain(url string) string {
	u := strings.Split(url, "//")[1]
	d := strings.Split(u, "/")[0]

	return d
}
