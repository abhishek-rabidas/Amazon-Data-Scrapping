package main

import (
	"bufio"
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"strings"
)

type Product struct {
	Name    string
	Price   string
	Link    string
	Rating  string
	Reviews string
	Image   string
}

var products []Product

// a-size-base-plus a-color-base a-text-normal

func main() {
	fmt.Println("Enter a product name:")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	c := colly.NewCollector()

	i := 0

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
		h.ForEach("div.a-section.a-spacing-base", func(_ int, h *colly.HTMLElement) {
			var name string
			name = h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
			var stars string
			stars = h.ChildText("span.a-icon-alt")
			var price string
			price = h.ChildText("span.a-price-whole")

			products = append(products, Product{Name: name, Rating: stars, Price: price})

		})
	})

	c.OnHTML(".a-size-base-plus.a-color-base.a-text-normal", func(e *colly.HTMLElement) {
		products = append(products, Product{Name: e.Text})
	})

	c.OnHTML(".a-text-price", func(e *colly.HTMLElement) {
		if i == len(products) {
			return
		}
		product := products[i]
		product.Price = e.Text
		products[i] = product
		i++
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit("https://amazon.in/s?k=" + strings.TrimRight(strings.ReplaceAll(line, " ", "+"), "\r\n"))
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Println("Name: ", product.Name)
		fmt.Println("Rating: ", product.Rating)
		fmt.Println("Price: ", product.Price)
		fmt.Println()
	}
}
