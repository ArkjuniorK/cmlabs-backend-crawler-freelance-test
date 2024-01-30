package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"path"
	"strings"

	"github.com/gocolly/colly"
)

var (
	url = flag.String("url", "https://telegram.org", "URL to crawl")
	wpt = ""
)

func main() {
	flag.Parse()
	log.Println("Initiate crawler")

	if len(*url) == 0 {
		log.Panicln(errors.New("crawler: empty url"))
	}

	domain := getDomain(*url)
	domains := getDomains(*url)

	wd, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}

	wpt = path.Join(wd, "result", domain)
	if err := handleDir(&wpt); err != nil {
		log.Panicln(err)
	}

	c := colly.NewCollector(
		colly.AllowedDomains(domains...),
	)

	start(c, url, &wpt)
}

func start(c *colly.Collector, url, wpt *string) {
	log.Println("Start crawling", *url)
	defer log.Println("Result saved at", *wpt)

	c.OnRequest(func(r *colly.Request) {
		up := r.URL.Path
		if len(up) == 0 {
			return
		}

		log.Println("Create dir for path", up)
		cp := path.Join(*wpt, up)
		if err := handleDir(&cp); err != nil {
			log.Println("An error has occured, aborting")
			log.Panicln(err)
		}
	})

	c.OnHTML("html", func(h *colly.HTMLElement) {
		cp := path.Join(*wpt, "index.html")

		up := h.Request.URL.Path
		if len(up) != 0 {
			p := path.Join(*wpt, h.Request.URL.Path)
			if err := handleDir(&p); err != nil {
				log.Println("An error has occured, aborting")
				log.Panicln(err)
			}

			cp = path.Join(*wpt, h.Request.URL.Path, "index.html")

			c.OnHTML("a[href]", func(h *colly.HTMLElement) {
				link := h.Attr("href")
				h.Request.Visit(link)
			})
			c.OnHTML("a[href]", func(h *colly.HTMLElement) {
				link := h.Attr("href")
				h.Request.Visit(link)
			})
		}

		if err := h.Response.Save(cp); err != nil {
			log.Println("An error has occured, aborting")
			log.Panicln(err)
		}
	})

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		h.Request.Visit(link)
	})

	c.OnError(func(rs *colly.Response, err error) {
		log.Println(
			"An error has occured for path",
			rs.Request.URL.Path,
			"with", err,
			"skipping!",
		)
	})

	c.Visit(*url)
}

func getDomain(url string) string {
	u := strings.Split(url, "//")[1]
	return strings.Split(u, "/")[0]
}

func getDomains(url string) []string {
	var (
		d       = getDomain(url)
		domains = make([]string, 0)
	)

	for _, v := range []string{"", "www."} {
		domains = append(domains, v+d)
	}

	return domains
}

func handleDir(p *string) error {
	err := os.MkdirAll(*p, os.ModePerm)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	if errors.Is(err, os.ErrExist) {
		if err = os.RemoveAll(*p); err != nil {
			return err
		}
	}

	return nil
}
