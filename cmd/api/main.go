package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Item struct {
	Name          string `json:"name"`
	Price         string `json:"price"`
	Coupon        string `json:"coupon"`
	ShippingPrice string `json:"shipping_price"`
}

func main() {
	/* Create a new collector and pass the domains he is allowed to scrape */
	c := colly.NewCollector(colly.AllowedDomains("listado.mercadolibre.com.co"))

	/*This method sets a callback function that will be executed when the specified
	HTML element is found during the scraping process. In this case, the callback */

	var items []Item

	c.OnHTML("div.poly-card__content", func(e *colly.HTMLElement) {
		item := Item{
			Name:          e.ChildText("h2.poly-box"),
			Price:         e.ChildText("div.poly-component__price div.poly-price__current span.andes-money-amount__fraction"),
			Coupon:        e.ChildText("div.poly-component__coupons div.poly-coupons__coupon-wrapper span.poly-coupons__coupon"),
			ShippingPrice: e.ChildText("div.poly-component__shipping"),
		}

		items = append(items, item)
	})

	/*
	* Not working good
	 */
	// c.OnHTML("[title]=Siguiente", func(e *colly.HTMLElement) {
	// 	next_page := e.Attr("href")
	// 	fmt.Println("Next page: ", next_page)
	// 	c.Visit(next_page)
	// })

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	/* Using the collector's visit function we can tell it to visit a specific url */
	c.Visit("https://listado.mercadolibre.com.co/juegos-de-mesa#D[A:juegos%20de%20mesa]")

	/* Convert the items to a JSON string */
	content, err := json.Marshal(items)

	if err != nil {
		fmt.Println("Error marshalling the items", err.Error())
	}

	os.WriteFile("result-scraping.json", content, 0644)

}
