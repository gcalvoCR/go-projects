package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Book struct{
  Title     string `json:"title"`
  Price     string `json:"price"`
  Status    string `json:"status"`
  Image    string `json:"image"`
}

func WebScraper(){

  startUrl := "https://books.toscrape.com/catalogue/page-1.html"

  books:= []Book{}
  csvBooks:= [][]string{}

  // Instance to work with library
  c := colly.NewCollector(
    // No need to add https or wwww
    colly.AllowedDomains("books.toscrape.com", "www.books.toscrape.com"),
  )

  // Find and visit all links
	c.OnHTML(".product_pod", func(e *colly.HTMLElement) {
		res := e.DOM

    title := res.Find("h3 a").Text()
    price := res.Find(".price_color").Text()
    status:= cleanData(res.Find(".instock").Text())
    image,_ := res.Find(".image_container a").Attr("href")

    book := Book{
      Title: title,
      Price: price,
      Status: status,
      Image: image,
    }
    csv := []string {title, price, status, image }

    books = append(books, book)
    csvBooks = append(csvBooks, csv)
	})

  // Continue method
	c.OnHTML(".next a", func(e *colly.HTMLElement) {
		res := e.DOM

    nextPage,_ := res.Attr("href")

    c.Visit(scrapeUrl(startUrl, nextPage))
	})


  c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

  c.OnError(func(_ *colly.Response, err error) {
    log.Println("Something went wrong:", err)
  })

  c.Visit(startUrl)

  writeJSON(books)
  writeCSV(csvBooks)

}

func writeCSV(data [][]string){
  csvFile, err := os.Create("books.csv")

if err != nil {
	log.Fatalf("failed creating file: %s", err)
}

  csvwriter := csv.NewWriter(csvFile)
  
  for _, empRow := range data {
    _ = csvwriter.Write(empRow)
  }
  csvwriter.Flush()
  csvFile.Close()
}

func writeJSON(data []Book){
  f, err := json.MarshalIndent(data, "", " ")

  if err!=nil{
    log.Fatal(err)
    return
  }

  _ = ioutil.WriteFile("books.json", f, 0644)
}

func main() {
  WebScraper()
}


func cleanData(data string) string {
  return strings.TrimSpace(data)
}

func scrapeUrl(startUrl, endpoint string) string{
  return strings.Replace(startUrl, "page-1.html", endpoint, 1)
}