package services

import (
	"Walker/app/utils"
	"Walker/app/utils/log"
	"Walker/global"
	"fmt"
	"github.com/gocolly/colly/v2"
)

type Reptile struct {
}

func (service *Reptile) Start() {
	c := colly.NewCollector(
		colly.Debugger(&log.ReptileLog{}),
		colly.AllowedDomains("m.woyaogexing.com"),
	)
	c.OnHTML("img[data-src]", func(e *colly.HTMLElement) {
		link := e.Attr("data-src")
		utils.DownLoad(global.BasePath+"/resources/", "https:"+link)
	})
	c.OnHTML(".m-page-2 li a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Visiting", e.Text)
		if e.Text == "下一条" {
			e.Request.Visit("https://m.woyaogexing.com" + link)
		}

	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	// 错误处理
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	// Start scraping on https://hackerspaces.org
	c.Visit("https://m.woyaogexing.com/touxiang/nv/2022/1235428.html")

	c.Wait()
}
